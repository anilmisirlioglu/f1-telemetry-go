const up = (str) => str.charAt(0).toUpperCase() + str.slice(1)

const out = (result, filename) => {
    switch(process.env.OUT){
        case 'TXT':
            require('fs').writeFileSync(`out-${filename}.txt`, result)
            break

        default:
            console.log(result)
    }
}

module.exports = {
    up,
    out
}