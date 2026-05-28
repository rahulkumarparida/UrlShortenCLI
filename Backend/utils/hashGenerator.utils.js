const crypto = require('crypto');

const generateHashId = () => {

    const unHashedToken = crypto.randomBytes(10).toString("hex")

    const hashedToken = crypto.createHash("sha256").update(unHashedToken).digest("hex").slice(0,7)

    return hashedToken

}

module.exports = {generateHashId}