const fs = require('fs')
const { out, up } = require('./utils')

const IN = process.env.IN
if(!IN){
    console.error('Please set `IN` env var.')
    process.exit(1)
}

const path = __dirname + `/data/${IN}.md`
if(!fs.existsSync(path)){
    console.error(`${IN}.md file not found.`)
    process.exit(1)
}

const content = fs.readFileSync(path).toString()
const data = content.split('|')
const arr = []
const PREFIX = process.env.PREFIX || ''
for(let i = 2; i < data.length; i = i + 3){
    const id = data[i - 1]
    const item = data[i]
    if(item && id){
        const formatGoConst = item.trim().split(' ').map(up).join('')
        arr.push(`${PREFIX}${formatGoConst} uint8 = ${id.trim()}`)
    }
}

out(arr.join('\n'), IN)