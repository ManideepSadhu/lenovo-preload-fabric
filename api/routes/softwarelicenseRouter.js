/**
 * Copyright 2018 IT People Corporation. All Rights Reserved.
 *
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an 'AS IS' BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */
'use strict';

const express = require('express');
const router = express.Router();
const verifyUserCanCall = require('../middleware/authorization');
const isAuthenticated = require('../middleware/authentication').authenticate;

let softwarelicense = require('../controllers/softwarelicenseController');

router.post('/softwarelicense', isAuthenticated, verifyUserCanCall, softwarelicense.recordSoftwareLicense);

router.put('/softwarelicense', isAuthenticated, verifyUserCanCall, softwarelicense.updateSoftwareLicense);

router.get('/softwarelicense/:License', isAuthenticated, verifyUserCanCall, softwarelicense.querySoftwareLicense);

router.get('/softwarelicensehistory/:License', isAuthenticated, verifyUserCanCall, softwarelicense.querySoftwareLicenseHistory);

router.get('/softwarelicenseowner/:Owner', isAuthenticated, verifyUserCanCall, softwarelicense.querySoftwareLicenseOwner);

router.get('/softwarelicenselist', isAuthenticated, verifyUserCanCall, softwarelicense.querySoftwareLicenseList);

router.delete('/softwarelicense/:License', isAuthenticated, verifyUserCanCall, softwarelicense.deleteSoftwareLicense);

module.exports = router;