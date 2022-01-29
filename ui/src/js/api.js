
/**
 * Class for handling HTTP interaction with Scoreboard API
 */
module.exports = class ScoreboardAPI {
    /**
     * @param {string} host HTTP endpoint hosting scoreboard API
     * @param {string?} jwt Token to authenticate with host
     * @todo Server/client auth
     */
    constructor(host = "http://localhost:8080") {
        this.host = host;
    }

    /**
     * Retrieve list of scoreboards
     * @returns {Promise<Object>} Map of scoreboard objects 
     */
    async list_scoreboards() {
        const url = `${this.host}/scores`;
        const response = await fetch(url, {method: 'GET'});
        const body = await response.json();
        return body.scoreboards;
    }

    /**
     * Given an ID, return that scoreboard
     * @param {string} id ID of scoreboard
     * @returns {Promise<Object?>} Returns scoreboard object if scoreboard exists, otherwise returns null
     */
    async get_scoreboard(id) {
        const url = `${this.host}/scores/${id}`;
        const params = { method: 'GET' };
        const response = await fetch(url, params);
        const body = await response.json();
        switch(response.status) {
            case 200:
                return body;
            case 404:
                return null;
            case 500:
                throw(body.message);
            default:
                console.error(response);
                throw(`Unsupported status code ${response.status}`);
        }
    }

    /**
     * Retrieve the featured scoreboard, or null if no scoreboard is currently featured
     * @returns {Object} Scoreboard object
     */
    async get_featured_scoreboard() {
        const url = `${this.host}/scores?featured=true`;
        const params = { method: 'GET' };
        const response = await fetch(url, params);
        const body = await response.json();
        switch(response.status) {
            case 200:
                return body;
            case 404:
                return null;
            case 500:
                throw(body.message);
            default:
                console.error(response);
                throw(`Unsupported status code ${response.status}`);
        }
    }

    /**
     * Create or update a scoreboard
     * @param {string} id ID of scoreboard
     * @param {Object} scoreboard Scoreboard object to pass to API
     */
    async set_scoreboard(id, scoreboard) {
        const url = `${this.host}/scores/${id}`;
        const params = {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(scoreboard)
        }
        const response = await fetch(url, params);
        if(response.status != 200) {
            const body = await response.json();
            throw(body.message);
        }
    }

    /**
     * Delete a scoreboard
     * @param {string} id ID of scoreboard
     */
    async remove_scoreboard(id) {
        const url = `${this.host}/scores/${id}`;
        const params = { method: 'DELETE' };
        const response = await fetch(url, params);
        if(response.status != 200) {
            const body = await response.json();
            throw(body.message);
        }
    }

    /**
     * Retrieve a list of current bracket IDs
     * @returns {string[]} List of available bracket IDs
     */
    async list_brackets() {
        const url = `${this.host}/brackets`;
        const params = { method: 'GET' };
        const response = await fetch(url, params);
        const body = await response.json();
        if(response.status === 200) {
            return body.brackets;
        } else if(response.status === 409) {
            window.alert(body.message);
        }
        throw(body.message);
    }
}