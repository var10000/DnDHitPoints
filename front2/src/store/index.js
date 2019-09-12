import Vue from 'vue';
import Vuex from 'vuex';
import createLogger from 'vuex/dist/logger';

import characters from './modules/characters';

import VuexPersist from 'vuex-persist';

const debug = process.env.NODE_ENV !== 'production';

// const vuexPersist = new VuexPersist({
//     key: 'my-app',
//     storage: sessionStorage // can be localStorage
// });

Vue.use(Vuex);

// const initialState = () => ({
//     ...characters.state,
// });
//
// const state = initialState();
//
// const getters = {
//     ...characters.getters,
// };
//
// const actions = {
//     ...characters.actions,
// };
//
// const mutations = {
//     ...characters.mutations,
// };

const plugins = [];

try {
    plugins.push(
        createLogger()
    );
    // plugins.push(
    //     vuexPersist.plugin
    // );
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