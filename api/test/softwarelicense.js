const axios = require("axios");
const expect = require("chai").expect;
const _ = require("lodash");
const nanoid = require('nanoid');
const log4js = require('log4js');
const logger = log4js.getLogger("SoftwareLicense-test");
describe("SoftwareLicense api ", async function() {
    this.timeout(45000);
    let xAccessToken;
    let payloadCreate = {
        "License": "1f2311b2-257f-423d-8e59-1efbdb65354e",
        "Owner": "elit ut veniam occaecat",
        "SoftwareName": "laborum enim culpa deserunt"
    };
    let payloadUpdate = {
        "License": "1f2311b2-257f-423d-8e59-1efbdb65354e",
        "Owner": "exercitation nostrud eu nulla",
        "SoftwareName": "consequat"
    };
    afterEach(function(done) {
        setTimeout(function() {
            done();
        }, 30000);
    });
    before(async () => {
        try {
            const email = `softwarelicense${nanoid()}@fakeemail.com;`
            await axios({
                method: 'post',
                url: "http://localhost:3000/users/register",
                data: {
                    email,
                    password: "p@ssw0rd",
                    org: "org1",
                    role: "admin"
                },
                headers: {
                    "content-type": "application/json"
                }
            })
            let response = await axios({
                method: 'post',
                url: "http://localhost:3000/users/login",
                data: {
                    email,
                    password: "p@ssw0rd",
                    org: "org1"
                },
                headers: {
                    "content-type": "application/json"
                }
            })
            xAccessToken = response.data.accessToken;
        } catch (error) {
            logger.error(error)
        }
    });

    it("Creates SoftwareLicense", async function() {
        const response = await axios({
            method: 'post',
            url: "http://localhost:3000/softwarelicense",
            data: payloadCreate,
            json: true,
            headers: {
                "content-type": "application/json",
                "x-access-token": xAccessToken
            }
        });
        payloadCreate.objectType = "SoftwareLicense";
        expect(response.status).to.be.equal(200);
        expect(payloadCreate).to.be.deep.equal(JSON.parse(response.data.objectBytes));
    });

    it("Updates SoftwareLicense", async function() {
        const response = await axios({
            method: 'put',
            url: "http://localhost:3000/softwarelicense",
            data: payloadUpdate,
            json: true,
            headers: {
                "content-type": "application/json",
                "x-access-token": xAccessToken
            }
        });
        payloadUpdate.objectType = "SoftwareLicense";
        expect(response.status).to.be.equal(200);
        expect(payloadUpdate).to.be.deep.equal(JSON.parse(response.data.objectBytes));
    });

    it("Gets a SoftwareLicense by id", async function() {
        const response = await axios({
            method: 'get',
            url: "http://localhost:3000/softwarelicense/1f2311b2-257f-423d-8e59-1efbdb65354e",
            headers: {
                "content-type": "application/json",
                "x-access-token": xAccessToken
            }
        })
        payloadUpdate.objectType = "SoftwareLicense";
        expect(response.status).to.be.equal(200);
        expect(payloadUpdate).to.be.deep.equal(JSON.parse(response.data.objectBytes));
    });

    it("Gets all SoftwareLicenses", async function() {
        const response = await axios({
            method: 'get',
            url: "http://localhost:3000/softwarelicenselist",
            headers: {
                "content-type": "application/json",
                "x-access-token": xAccessToken
            }
        });
        expect(response.status).to.be.equal(200);
        expect(JSON.parse(response.data.objectBytes)).to.deep.include(payloadUpdate);

    });

    it("Deletes SoftwareLicense", async function() {
        const response = await axios({
            method: 'delete',
            url: "http://localhost:3000/softwarelicense/1f2311b2-257f-423d-8e59-1efbdb65354e",
            headers: {
                "content-type": "application/json",
                "x-access-token": xAccessToken
            }
        });
        expect(response.status).to.be.equal(200);
        expect(response.data.objectBytes).to.be.equal('');
    });
});