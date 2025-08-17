const AssistantMessage = 1
const UserMessage = 2

const AssistantName = 'burp'
const StatusName = '!status'

class Chat {
  constructor({
    nickname = 'anon',
    channel = 'status',
    model,
    temperature,
    maxTokens,
    topP,
    topK,
    subscribeUrl,
    publishUrl,
  } = {}) {
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

    this.temperature = temperature
    this.maxTokens = maxTokens
    this.topP = topP
    this.topK = topK
  }

  setUserNickname(nickname = this.nickname) {
    nickname = sanitizeAlphaNum(nickname, 32)
    if (!nickname) {
      return
    }
    this.nickname = nickname
    if (this.elements.userLabel) {
      this.elements.userLabel.textContent = `[${this.nickname}/${this.model}]`
    }
  }

  setChannelName(name = this.channel) {
    name = sanitizeAlphaNum(name, 32)
    if (!name) {
      return
    }
    this.channel = name
    if (this.elements.channelLabel) {
      this.elements.channelLabel.textContent = `[${(this.channel === 'status' ? '!' : '#') + this.channel}]`
    }
  }

  setParams(temperature = this.temperature, maxTokens = this.maxTokens, topP = this.topP, topK = this.topK) {
    // TODO: validate
    this.temperature = temperature
    this.maxTokens = maxTokens
    this.topP = topP
    this.topK = topK

    if (this.elements.paramsLabel) {
      const container = this.elements.paramsLabel

      // Clear out old contents
      container.innerHTML = ''

      const addParam = (emoji, value, label, raw = null) => {
        const span = document.createElement('span')
        span.setAttribute('title', `${label} ${raw || value}`)
        span.textContent = `${emoji} ${value}`
        container.appendChild(span)
      }

      addParam('🔥', parseFloat(this.temperature.toFixed(7)), 'temperature')

      if (Number.isFinite(this.topP)) {
        addParam('🔮', parseFloat(this.topP.toFixed(7)), 'top-p')
      }

      if (Number.isInteger(this.topK)) {
        addParam('🔑', this.topK, 'top-k')
      }

      if (Number.isInteger(this.maxTokens)) {
        addParam('⏳', formatMaxTokens(this.maxTokens), 'max tokens', this.maxTokens)
      }
    }
  }

  setParams1(temperature = this.temperature, maxTokens = this.maxTokens, topP = this.topP, topK = this.topK) {
    // TODO: validate
    this.temperature = temperature
    this.maxTokens = maxTokens
    this.topP = topP
    this.topK = topK
    if (this.elements.paramsLabel) {
      let s = '🔥 ' + parseFloat(this.temperature.toFixed(7))
      if (Number.isFinite(this.topP)) {
        s += ' 🔮 ' + parseFloat(this.topP.toFixed(7))
      }
      if (Number.isInteger(this.topK)) {
        s += ' 🔢 ' + this.topK
      }
      if (Number.isInteger(this.maxTokens)) {
        s += ' ⏳ ' + formatMaxTokens(this.maxTokens)
      }
      this.elements.paramsLabel.textContent = `[${s}]`
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
      paramsLabel: root.getElementById('params'),
      spinner: root.getElementById('spinner'),
    }

    this.setUserNickname()
    this.setChannelName()
    this.setParams()

    this.setupClock()
    this.bindInput()
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
    } else if (sender === StatusName) {
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

  bindInput() {
    this.elements.form.addEventListener('submit', e => {
      e.preventDefault()
      const msg = this.elements.input.value.trim()
      if (!msg) {
        return
      }

      this.elements.input.value = ''

      this.startSpinner()

      this._send(msg)
        .then(res => {
          if (!res.ok) {
            this.stopSpinner()
            this.addMessage(
              ['failed to send message:', res.statusText.toLowerCase(), String(res.status)].join(' '),
              StatusName,
              new Date(),
            )
          }
        })
        .catch(err => {
          this.stopSpinner()
          this.addMessage(
            'failed to send message: ' + (err instanceof Error ? err.message : String(err)).toLowerCase(),
            StatusName,
            new Date(),
          )
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
        this.addMessage(
          ['failed to retrieve chat history:', res.statusText.toLowerCase(), String(res.status)].join(' '),
          StatusName,
          new Date(),
        )
        return
      }
      const messages = await res.json()
      messages
        .slice()
        .reverse()
        .forEach(msg => this.renderMessage(msg))
    } catch (e) {
      this.addMessage(
        'failed to retrieve chat history: ' + (err instanceof Error ? err.message : String(err)).toLowerCase(),
        StatusName,
        new Date(),
      )
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

        const msg = await res.json()
        this.renderMessage(msg)
      } catch (err) {
        this.addMessage(
          'poll error: ' + (err instanceof Error ? err.message : String(err)).toLowerCase() + '; will retry in 5s',
          StatusName,
          new Date(),
        )
        await new Promise(r => setTimeout(r, 5000))
      }
    }
  }

  // Protected methods to be overridden
  async _send(msg) {
    const u = new URL('/ask', this.publishUrl)
    u.searchParams.set('id', this.channel)
    u.searchParams.set('model', this.model)
    u.searchParams.set('temp', this.temperature)
    u.searchParams.set('max_tokens', this.maxTokens)
    if (Number.isFinite(this.topP)) {
      u.searchParams.set('top_p', this.topP)
    }
    if (Number.isInteger(this.topK)) {
      u.searchParams.set('top_k', this.topK)
    }
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

function formatMaxTokens(n) {
  if (n >= 1_000_000) {
    return (n / 1_000_000).toFixed(1).replace(/\.0$/, '') + 'M'
  }
  if (n >= 1_000) {
    return (n / 1_000).toFixed(1).replace(/\.0$/, '') + 'k'
  }
  return n.toString()
}
