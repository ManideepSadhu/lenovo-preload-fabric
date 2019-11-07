
"use strict";

const express = require('express');
const router = express.Router();



     
const softwarelicenseRoute = require('./softwarelicenseRouter.js');
router.use(softwarelicenseRoute);
    


module.exports = router;