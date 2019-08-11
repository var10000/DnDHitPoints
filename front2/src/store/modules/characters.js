import axios from "axios";
import Vue from "vue";
import { format } from "date-fns";

const characterListURL = `http://172.20.10.4:10000/rooms/15/characterList`;

const initialState = () => ({
    listOfCharacters: [],
    listOfFightCharacters: []
});

const state = initialState();

const getters = {
    getListOfCharacters: ({ listOfCharacters }) => listOfCharacters,
    getListOfFightCharacters: ({ listOfFightCharacters }) => listOfFightCharacters,
};

const actions = {
    loadListOfCharacters({ commit }) {
        const error_statuses = [401, 404, 403];

        return axios
            .get(characterListURL)
            .then(({ data: characterList }) => {
                console.log(characterList);
                commit("loadedListOfCharacters", characterList);
            })
            .catch(error => {
                if (
                    error.response && error_statuses.includes(error.response.status)
                ) {
                    throw error;
                }
                alert(error);
            });
    },
    sendScheduleItem({ commit, getters }, { data, date }) {

        const axiosConfig = {
            headers: getters["profile/authPostHeader"]
        };

        const meetup_code = getters["profile/getMeetupCode"];

        return axios
            .post(`${meetupURL}${meetup_code}/schedule/`, data, axiosConfig)
            .then(({ data: schedule_item }) => {
                commit("sendedScheduleItem", { schedule_item, date });
            })
            .catch(error => {
                throw error;
            });
    },

    deleteScheduleItem({ commit, getters }, { item_id }) {
        const axiosConfig = {
            headers: getters["profile/authPostHeader"]
        };

        return axios
            .delete(
                `${scheduleURL}${item_id}/`,
                axiosConfig
            )
            .then(() => {
                commit("deletedScheduleItem", {
                    item_id
                });
            })
            .catch(error => {
                if (error.response && error.response.status === 401) {
                    return false;
                }
                versionChecker(error, () => commit("profile/logout"));
            });
    },
};

const mutations = {
    loadedListOfCharacters(state, characterList) {
        console.log(Array.from(characterList));
        state.listOfCharacters = Array.from(characterList);
    },
    sendedScheduleItem(state, { schedule_item, date }) {
        let group = state.schedule.find(item => item.id == schedule_item.group);
        if (!group) {
            Vue.set(state.schedule, state.schedule.length, {
                id: schedule_item.group,
                date: date,
                items: []
            });
        }

        let activity_id_to_send = (schedule_item.activity_id) ? schedule_item.activity_id.activity_id : null;
        // rework data because not flat from back
        const dataToInsert = {
            ...schedule_item,
            activity_id: activity_id_to_send
        };

        console.log(state.activities);
        console.log(state.events);

        state.activities[activity_id_to_send] = {...schedule_item, ...schedule_item.activity_id, questions: [], materials: [], polls: [] };

        // schedule item to events
        state.events[dataToInsert.id] = dataToInsert;

        // schedule item to schedule
        state.schedule.find(item => item.id == schedule_item.group).items.push(schedule_item.id);

        state.schedule.sort((a, b) => {
            if (a.date < b.date) return -1;
            return 1;
        });

    },
    deletedListCharacterItem(state, { item_id }) {

        const group_id = state.events[item_id].group;
        const group = state.schedule.find(item => item.id == group_id);

        Vue.delete(state.events, item_id);

        //filter items array of days and delete day if no items
        state.schedule.find(item => item.id == group_id).items = group.items.filter(
            item => item != item_id
        );
        if (!group.items.length) {
            state.schedule = state.schedule.filter(item => item.id != group.id);
        }
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};

