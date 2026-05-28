const express = require('express');
const {urlRegister , getUrl , redierectUrl, urlAnalytics } = require('../controllers/url.controller.js');


const urlRouter = express.Router()
const redierectRouter = express.Router()

urlRouter.route('/url').get(getUrl).post(urlRegister)
redierectRouter.route('/:code').get(redierectUrl)
redierectRouter.route('/analytics/:code').get(urlAnalytics)

module.exports = {urlRouter ,redierectRouter}