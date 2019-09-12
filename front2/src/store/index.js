import Vue from 'vue';
import Vuex from 'vuex';
import createLogger from 'vuex/dist/logger';

import characters from './modules/characters';

const debug = process.env.NODE_ENV !== 'production';

Vue.use(Vuex);

const plugins = [];

try {
    plugins.push(
        createLogger()
    );
} catch (e) {
    console.error(e);
}

export default new Vuex.Store({
    modules: {
        characters
    },
    strict: debug,
    plugins
});