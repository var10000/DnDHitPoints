<template>
    <!-- Popup -->
    <f7-popup
            id="create-battle-list-popup"
            :opened="opened"
            @popup:close="$emit('popup:closed')">
        <f7-page>
            <f7-navbar class="navbar-panel">
                <f7-nav-left>
                    <f7-link
                            popup-close
                            class="navbar-panel__text"
                            color="white"
                            text="Закрыть" />
                </f7-nav-left>
                <f7-nav-subtitle class="navbar-panel__title"> Creating </f7-nav-subtitle>
                <f7-nav-right>
                    <f7-link
                            @click="complete"
                            class="navbar-panel__text"
                            color="white"
                            text="Сохранить" />
                </f7-nav-right>
            </f7-navbar>

            <f7-block >
                <!-- Personal data -->
                <f7-block-title>Выберите персонажа</f7-block-title>
                    <div v-if="listOfCharacters.length">
                        <div
                                v-for="character in listOfCharacters"
                                :key="character.ID">

                            <p>{{ character.ID}}</p>
                        </div>
                    </div>
            </f7-block>
        </f7-page>
    </f7-popup>
</template>

<script>
    export default {
        name: "CreateBattleListPopup",
        props: {
            opened: {
                type: Boolean,
                default: false
            },
            listOfCharactersData: {
                type: Array,
                default: []
            },
            listOfFightCharactersData: {
                type: Array,
                default: []
            }
        },
        data() {
            return {
                // name: "",
                // armorType: "",
                // initiative: null,
                // hits: null,
                listOfCharacters: this.listOfCharactersData || [],
                listOfFightCharacters: this.listOfFightCharactersData || []
            }
        },
        watch: {
            listOfCharactersData(v) {
                console.log(v);
                v.forEach((item) => console.log(item));
                this.listOfCharacters = v;
            }
        },
        methods: {
            complete() {
                let data = {
                    name: this.name,
                    armorType: this.armorType,
                    initiative: this.initiative,
                    hits: this.hits,
                }
                this.$emit("popup:complete", data);
                this.$emit("popup:closed");
            }
        },
    }
</script>

<style src="../css/popup.css"></style>

<!--<template>-->
    <!--&lt;!&ndash; Popup &ndash;&gt;-->
    <!--<f7-popup-->
            <!--id="create-battle-list-popup"-->
            <!--:opened="opened"-->
              <!--@popup:close="$emit('popup:closed')">-->
        <!--<f7-page>-->
            <!--<f7-navbar class="navbar-panel">-->
                <!--<f7-nav-left>-->
                    <!--<f7-link-->
                            <!--popup-close-->
                            <!--class="navbar-panel__text"-->
                            <!--color="white"-->
                            <!--text="Закрыть" />-->
                <!--</f7-nav-left>-->
                <!--<f7-nav-subtitle class="navbar-panel__title"> Creating </f7-nav-subtitle>-->
                <!--<f7-nav-right>-->
                    <!--<f7-link-->
                            <!--@click="complete"-->
                            <!--class="navbar-panel__text"-->
                            <!--color="white"-->
                            <!--text="Сохранить" />-->
                <!--</f7-nav-right>-->
            <!--</f7-navbar>-->
        </f7-page>
        <f7-page-content>
            <!-- List -->
            <div v-if="listOfCharacters.length">
                <div
                        v-for="character in listOfCharacters"
                        :key="character.id">
                        <p>{{ character }}</p>
                        <!--<activity-item-->
                                <!--v-for="activity in day.items"-->
                                <!--v-if="activity"-->
                                <!--:class="{ 'themed-schedule': !isNotThemed }"-->
                                <!--:href="`/event/${day.id}/${activity.activity_id}/${activity.id}/${exactHref(activity.tab_questions, activity.tab_polls, activity.tab_materials)}`"-->
                                <!--:key="activity.id"-->
                                <!--:schedule-id="day.id"-->
                                <!--:activity="activity"-->
                                <!--:date="day.date"-->
                                <!--:is-vote-loading="isVoteLoading(activity.id)"-->
                                <!--:with-time="isNotThemed"-->
                                <!--:new-view="!isNotThemed"-->
                                <!--@refresh="refresh"-->
                                <!--@voteFor="voteFor"-->
                                <!--@voteAgainst="voteAgainst"-->
                        <!--/>-->
                </div>
            </div>
        <!--</f7-page-content>-->
    <!--</f7-popup>-->
<!--</template>-->

<!--<script>-->
    <!--export default {-->
        <!--name: "CreateBattleListPopup",-->
        <!--props: {-->
            <!--opened: {-->
                <!--type: Boolean,-->
                <!--default: false-->
            <!--},-->
            <!--listOfCharactersData: {-->
                <!--type: Array,-->
                <!--default: []-->
            <!--},-->
            <!--listOfFightCharactersData: {-->
                <!--type: Array,-->
                <!--default: []-->
            <!--}-->
        <!--},-->
        <!--data() {-->
            <!--return {-->
                <!--// name: "",-->
                <!--// armorType: "",-->
                <!--// initiative: null,-->
                <!--// hits: null,-->
                <!--listOfCharacters: this.listOfCharactersData || [],-->
                <!--listOfFightCharacters: this.listOfFightCharactersData || []-->
            <!--}-->
        <!--},-->
        <!--methods: {-->
            <!--complete() {-->
                <!--console.log(this.listOfCharacters);-->
                <!--// let data = {-->
                <!--//     name: this.name,-->
                <!--//     armorType: this.armorType,-->
                <!--//     initiative: this.initiative,-->
                <!--//     hits: this.hits,-->
                <!--// }-->
                <!--this.$emit("popup:complete");-->
                <!--this.$emit("popup:closed");-->
            <!--}-->
        <!--}-->
    <!--}-->
<!--</script>-->

<!--<style src="../css/popup.css"></style>-->