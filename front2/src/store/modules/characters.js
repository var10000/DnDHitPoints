import axios from "axios";
import Vue from "vue";
import { format } from "date-fns";
import { ru } from "date-fns/locale";


import versionChecker from "../versionChecker";
//
// const meetupURL = `http://${process.env.VUE_APP_API_HOST}:${
//     process.env.VUE_APP_API_PORT
//     }/api/meetup/`;
const scheduleURL = `http://${process.env.VUE_APP_API_HOST}:${
    process.env.VUE_APP_API_PORT
    }/api/schedule/`;

const initialState = () => ({
    listOfCharacters: ["1", "2", "3"],
    listOfFightCharacters: []
});

const state = initialState();

const getters = {
    getListOfCharacters: ({ listOfCharacters }) => listOfCharacters,
    getListOfFightCharacters: ({ listOfFightCharacters }) => listOfFightCharacters,
};

const actions = {
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
    editScheduleItem({ commit, getters }, { event_id, activity_id, difference, date }) {
        const axiosConfig = {
            headers: getters["profile/authPostHeader"]
        };
        difference.group = 47;
        let schedule_item = difference;
        commit("editedScheduleItem", { event_id, activity_id, schedule_item, date });
        // return axios
        //   .patch(`${scheduleURL}${event_id}/`, difference, axiosConfig)
        //   .then(({ data: difference }) => {
        //     commit("editedScheduleItem", { event_id, activity_id, difference, date });
        //   })
        //   .then(() => {
        //     commit("editedScheduleItem", { event_id, activity_id, difference, date });
        //   })
        //   .catch(error => {
        //     if (error.response && error.response.status === 304) {
        //       throw Error(304);
        //     }
        //   });
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
    exportData({ getters }, { meetup_code, data }) {

        const axiosConfig = {
            headers: getters["profile/authHeader"],
            responseType: "arraybuffer"
        };

        const exportQuestionsURL =
            `${meetupURL}${meetup_code}/questions/export/?format=${data.questions.format}&comments=${data.questions.comments}&likes=${data.questions.likes}`;

        return axios
            .get(exportQuestionsURL, axiosConfig)
            .then(response => {

                const blob = new Blob([response.data]);

                let link = document.createElement("a");
                const now = format(new Date(), 'dd(ccc)_MM_yyyy-HH_mm_ss', { locale: ru });
                const filename = `questions_${meetup_code}-${now}`;
                link.href = window.URL.createObjectURL(blob);
                link.download = `${filename}.${data.questions.format}`;
                link.click();
                window.URL.revokeObjectURL(link);
            })
            .catch(error => {
                if (error.response && error.response.status) {
                    throw error;
                }
            })
    }

};

const mutations = {

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
    editedScheduleItem(state, { event_id, activity_id, schedule_item, date }) {
        console.log(schedule_item);
        const { theme, speaker, start_time, end_time, group, activity } = schedule_item;
        console.log(activity);
        console.log(activity.tab_questions);
        if (theme) Vue.set(state.events[event_id], "theme", theme);
        if (speaker) Vue.set(state.events[event_id], "speaker", speaker);
        if (start_time) Vue.set(state.events[event_id], "start_time", start_time);
        if (end_time) Vue.set(state.events[event_id], "end_time", end_time);

        let day = state.schedule.findIndex(day => (day.items.find(item => item === event_id)));
        console.log("day");
        console.log(day);
        let item = state.schedule[day].items.findIndex(item => item === event_id);
        console.log("item");
        console.log(item);
        state.schedule[day].items.splice(item, 1);

        console.log("In MAIN STORE");
        if (group) {
            console.log("i am group");
            console.log(group);
            let dayExists = state.schedule.find(item => item.id == group);
            console.log("i am day");
            console.log(dayExists);
            if (!dayExists) {
                console.log("budet set");
                Vue.set(state.schedule, state.schedule.length, {
                    id: group,
                    date: date,
                    items: []
                });
            }
        }

        // rework data because not flat from back
        const dataToInsert = {
            ...schedule_item,
            activity_id: activity_id
        };

        // schedule item to events
        state.events[event_id].group = group;
        //state.schedule.forEach(day => (day.items.filter(item => item !== event_id)));
        //console.log(state.schedule[index].items.find(item => item === event_id));


        // schedule item to schedule
        state.schedule.find(item => item.id === group).items.push(event_id);

        console.log("schedule");
        console.log(state.schedule);
        console.log("events");
        console.log(state.events);

        state.schedule.sort((a, b) => {
            if (a.date < b.date) return -1;
            return 1;
        });

        if (activity.tab_questions !== undefined) Vue.set(state.activities[activity_id], "tab_questions", activity.tab_questions);
        if (activity.tab_polls !== undefined) Vue.set(state.activities[activity_id], "tab_polls", activity.tab_polls);
        if (activity.tab_materials !== undefined) Vue.set(state.activities[activity_id], "tab_materials", activity.tab_materials);
        if (activity.allow_questions !== undefined) Vue.set(state.activities[activity_id], "allow_questions", activity.allow_questions);

        console.log("activities");
        console.log(state.activities);

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
    state,
    getters,
    actions,
    mutations
};

