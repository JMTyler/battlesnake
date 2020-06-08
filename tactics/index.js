const _  = require('lodash');
const fs = require('fs');

const files = fs.readdirSync(__dirname).filter((file) => (file !== 'index.js'));
const tactics = _.map(files, (file) => require(`./${file}`));

module.exports = Object.assign({}, ...tactics);