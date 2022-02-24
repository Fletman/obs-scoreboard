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
     * @param {string[]?} score_ids Optional list of score-ids to filter in request
     * @returns {Promise<Object>} Map of scoreboard objects 
     */
    async list_scoreboards(score_ids = null) {
        const url = score_ids && score_ids.length > 0 ?
            `${this.host}/scores?${score_ids.map((id) => `score-id=${id}`).join('&')}`:
            `${this.host}/scores`;
        const params = { method: 'GET' };
        const response = await fetch(url, params);
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
     * @returns {Promise<Object>} Scoreboard object
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
        };
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
     * @returns {Promise<string[]>} List of available bracket IDs
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

    /**
     * Given an ID, return that bracket
     * @param {string} id ID of bracket
     * @returns {Promise<Object>}
     */
    async get_bracket(id) {
        const url = `${this.host}/brackets/${id}`;
        const params = { method: 'GET' };
        const response = await fetch(url, params);
        const body = await response.json();
        if(response.status != 200) {
            throw(body.message);
        } else {
            return body;
        }
    }

    /**
     * Create a new bracket
     * @param {string} id ID of bracket
     * @param {string[]} seeds Array of team names, prdered by seeding
     * @param {Number} match_size Size of each match, defaults to 2
     * @returns {Promise<Object>}
     */
    async create_bracket(id, seeds, match_size = 2) {
        const url = `${this.host}/brackets`;
        const request = { 'bracket-id': id, 'match-size': match_size, teams: seeds };
        const params = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(request)
        };
        const response = await fetch(url, params);
        const body = await response.json();
        if(response.status != 200) {
            throw(body.message);
        } else {
            return body;
        }
    }
}