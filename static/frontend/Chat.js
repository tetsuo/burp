const AssistantMessage = 1
const UserMessage = 2

const AssistantName = 'burp'
const HelpName = 'help'

class Chat {
  constructor({ nickname = 'anon', channel = 'status', model, subscribeUrl, publishUrl } = {}) {
    // alphanumeric only, max 32 chars
    this.nickname = sanitizeAlphaNum(nickname, 32) || 'anon'
    this.channel = sanitizeAlphaNum(channel, 32) || 'status'

    this.subscribeUrl = subscribeUrl
    this.publishUrl = publishUrl

    this.lastTime = null
    this.msgBuffer = ''
    this.elements = {}
    this.spinner = { timer: null, i: 0, running: false, frames: ['-', '\\', '|', '/'] }
    this.model = model
  }

  setUserNickname(nickname) {
    if (this.elements.userLabel) {
      nickname = sanitizeAlphaNum(nickname, 32)
      if (!nickname) {
        return
      }
      this.nickname = nickname
      this.elements.userLabel.textContent = `[${this.nickname}/${this.model}]`
    }
  }

  setChannelName(name) {
    if (this.elements.channelLabel) {
      name = sanitizeAlphaNum(name, 32)
      if (!name) {
        return
      }
      this.channel = name
      this.elements.channelLabel.textContent = `[${(this.channel === 'status' ? '!' : '#') + this.channel}]`
    }
  }

  mountTo(root = document) {
    this.elements = {
      chatLines: root.getElementById('chat-lines'),
      input: root.getElementById('chat-input'),
      form: root.getElementById('chat-form'),
      userLabel: root.getElementById('user'),
      channelLabel: root.getElementById('channel'),
      timeLabel: root.getElementById('time'),
      spinner: root.getElementById('spinner'),
    }

    this.setUserNickname(this.nickname)
    this.setChannelName(this.channel)

    this.setupClock()
    this.bindForm()
    this.loadRecent().then(() => this.pollLoop())
  }

  startSpinner() {
    if (this.spinner.running || !this.elements.spinner) {
      return
    }
    this.spinner.running = true
    this.elements.spinner.style.visibility = 'visible'
    this.elements.spinner.textContent = this.spinner.frames[this.spinner.i]

    this.spinner.timer = setInterval(() => {
      this.spinner.i = (this.spinner.i + 1) % this.spinner.frames.length
      this.elements.spinner.textContent = this.spinner.frames[this.spinner.i]
    }, 130)
  }

  stopSpinner() {
    if (!this.spinner.running) {
      return
    }
    clearInterval(this.spinner.timer)
    this.spinner.timer = null
    this.spinner.running = false
    if (this.elements.spinner) {
      this.elements.spinner.textContent = ''
      this.elements.spinner.style.visibility = 'hidden'
    }
  }

  setupClock() {
    const tick = () => {
      if (this.elements.timeLabel) {
        this.elements.timeLabel.textContent = `[${nowHHMMSS()}]`
      }
      setTimeout(tick, 1000)
    }
    tick()
  }

  addMessage(text, sender, timeISO) {
    if (!text || !text.trim()) {
      return
    }

    const line = document.createElement('div')
    line.className = 'line'

    const tspan = document.createElement('span')
    tspan.className = 'time'
    const stamp = timeISO ? (timeISO instanceof Date ? timeISO : new Date(timeISO)) : new Date(this.lastTime)
    tspan.textContent = `[${pad(stamp.getHours())}:${pad(stamp.getMinutes())}:${pad(stamp.getSeconds())}]`

    const uspan = document.createElement('span')
    uspan.className = 'user'
    if (sender === AssistantName) {
      uspan.className = 'assistant'
    } else if (sender === HelpName) {
      uspan.className = 'help'
    }
    uspan.textContent = `<${sender}>`

    const mspan = document.createElement('span')
    mspan.className = 'message'

    const parts = text.split(/((?:https?:|magnet:|\/ipfs\/)\\S+)/g)
    parts.forEach(p => {
      if (/^(?:https?:|magnet:|\/ipfs\/)/.test(p)) {
        const a = document.createElement('a')
        a.href = p
        a.textContent = p
        a.target = '_blank'
        a.rel = 'noreferrer noopener nofollow ugc'
        mspan.appendChild(a)
      } else {
        mspan.appendChild(document.createTextNode(p))
      }
    })

    line.append(tspan, uspan, mspan)
    this.elements.chatLines.appendChild(line)

    this.elements.chatLines.scrollTop = this.elements.chatLines.scrollHeight
  }

  bindForm() {
    this.elements.form.addEventListener('submit', e => {
      e.preventDefault()
      const msg = this.elements.input.value.trim()
      if (!msg) {
        return
      }

      this.elements.input.value = ''

      this.startSpinner()

      this._send(msg).catch(err => {
        console.error('send failed:', err)
        this.stopSpinner()
        this.addMessage('try again; send failed: ' + err, HelpName, new Date())
      })
    })
  }

  updateCursor(iso) {
    if (iso) {
      this.lastTime = iso
    }
  }

  renderMessage(msg) {
    // Track cursor from server for correct long-poll resume
    this.updateCursor(msg.Time)

    // Then ignore timeouts
    if (msg.LongPollTimeout) {
      return
    }

    const body = msg.Body || ''
    if (msg.Role !== AssistantMessage && !body.trim()) {
      return
    }

    // Handle by role
    if (msg.Role === UserMessage) {
      this.addMessage(body, this.nickname, msg.Time)
      return
    }

    if (msg.Role === AssistantMessage) {
      // EOF:
      if (body === '') {
        // flush any remaining buffered text
        const msg = this.msgBuffer.trim()
        this.msgBuffer = ''
        if (msg) {
          this._recv(msg).catch(err => {
            this.addMessage(err.message || String(err), AssistantName)
          })
        }
        this.stopSpinner()
        return
      }

      // Collapse runs of whitespace + newlines into a single newline
      const normalized = body.replace(/[ \t]*(?:\r?\n[ \t]*)+/g, '\n')

      // Append to buffer and split into lines
      this.msgBuffer += normalized
      const lines = this.msgBuffer.split(/\r?\n/)
      this.msgBuffer = lines.pop() // keep unfinished line buffered

      for (const line of lines) {
        const text = line.trim()
        if (!text) {
          continue
        }
        this.addMessage(text, AssistantName)
      }
      return
    }

    console.error(`unknown role "${msg.Role}"; skipping message: ${JSON.stringify(msg)}`)
  }

  async loadRecent() {
    try {
      const u = new URL('/recent', this.subscribeUrl)
      u.searchParams.set('id', this.channel)

      const res = await fetch(u.toString())
      if (!res.ok) {
        throw new Error(res.status)
      }
      const messages = await res.json()
      messages
        .slice()
        .reverse()
        .forEach(msg => this.renderMessage(msg))
    } catch (e) {
      console.error('initial load failed:', e)
    }
  }

  async pollLoop() {
    while (true) {
      try {
        const u = new URL('/wait', this.subscribeUrl)
        u.searchParams.set('id', this.channel)
        if (this.lastTime) {
          u.searchParams.set('after', this.lastTime)
        }

        const res = await fetch(u.toString())
        if (!res.ok) {
          throw new Error('status ' + res.status)
        }

        const ev = await res.json()
        this.renderMessage(ev)
      } catch (err) {
        console.error('poll error; will retry in 5s', err)
        await new Promise(r => setTimeout(r, 5000))
      }
    }
  }

  // Protected methods to be overridden
  async _send(msg) {
    const u = new URL('/ask', this.publishUrl)
    u.searchParams.set('id', this.channel)
    u.searchParams.set('model', this.model)
    return fetch(u.toString(), {
      method: 'POST',
      headers: { 'Content-Type': 'text/plain' },
      body: msg,
    })
  }

  async _recv(msg) {
    this.addMessage(msg, AssistantName)
  }
}

function pad(n) {
  return n < 10 ? '0' + n : n
}

function nowHHMMSS() {
  const d = new Date()
  return `${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`
}

function sanitizeAlphaNum(value, maxLen = 32) {
  const s = String(value ?? '').replace(/[^0-9A-Za-z]/g, '')
  return s.slice(0, maxLen)
}
