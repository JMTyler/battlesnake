const _    = require('lodash');
const fs   = require('fs');
const path = require('path');

const thisFile = path.basename(__filename);
const files = fs.readdirSync(__dirname).filter((file) => (file !== thisFile));
const tactics = _.map(files, (file) => require(`./${file}`));

module.exports = Object.assign({}, ...tactics);
