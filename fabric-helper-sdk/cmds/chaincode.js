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
 * Author: Sandeep Pulluru <sandeep.pulluru@itpeoplecorp.com>
 */
'use strict';

exports.command = 'chaincode <command>';
exports.desc = 'Install, instantiate and update chaincode';
exports.builder = function (yargs) {
    return yargs.commandDir('chaincode-cmds')
        .option('org', {
            demandOption: true,
            describe: 'Name of the org',
            requiresArg: true,
            type: 'string'
        })
};