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
                            <div v-for="character in listOfCharacters"
                                 :key="character.ID">
                                <!--<p>{{character}}</p>-->
                                <f7-button v-on:click="toFightList(character)"  style="background-color: #f0e6d5; margin: 5px; min-height: 140px; color: saddlebrown">
                                    <div>
                                        <h3> {{ character.Name }} id: {{ character.ID }} <br>
                                            ArmorType: {{ character.ArmorType }} <br>
                                            Initiative: {{ character.Initiative }} <br>
                                            Hits: {{ character.Hits }}</h3>

                                    </div>
                                </f7-button>
                                <!--<character-card-->
                                        <!--:character-data="character"-->
                                <!--v-on:click="toFightList">-->
                                <!--</character-card>-->
                            </div>
                    </div>
            </f7-block>
        </f7-page>
    </f7-popup>
</template>

<script>
    import CharacterCard from "./CharacterCard"
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
        components:  {
            CharacterCard
        },
        data() {
            return {
                listOfCharacters: this.listOfCharactersData.slice() || [],
                listOfFightCharacters: this.listOfFightCharactersData.slice() || []
            }
        },
        watch: {
            listOfCharactersData(v) {
                this.listOfCharacters = v;
            }
        },
        methods: {
            toFightList(character) {
                let index2 = this.listOfCharacters.findIndex((char) => char.ID === character.ID);
                this.listOfCharacters = this.listOfCharacters.splice(index2,1);
                this.listOfFightCharacters.push(character);
                // let index = this.listOfFightCharacters.findIndex((char) => char.ID === character.ID);
                // if (index === -1)  {
                //     this.listOfFightCharacters.push(character);
                //     let index2 = this.listOfCharacters.findIndex((char) => char.ID === character.ID);
                //     this.listOfCharacters = this.listOfCharacters.splice(index2,1);
                //     console.log("remove");
                //     console.log(this.listOfCharacters);
                // } else {
                //     this.listOfFightCharacters = this.listOfFightCharacters.splice(index,1);
                //     this.listOfCharacters.push(character);
                //     console.log("add");
                //     console.log(this.listOfCharacters);
                // }
            },
            complete() {
                let listOfFightCharacters = this.listOfFightCharacters;
                this.$emit("popup:complete", listOfFightCharacters);
                this.$emit("popup:closed");
            }
        },
    }
</script>

<style src="../css/popup.css"></style>