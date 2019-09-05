<template>
  <f7-app id="app">
    <div class="room">
      <header class="header">
        <div class="header__text">
          <p class="header__text-title"> D&D craft page</p>
          <p class="header__text-title_description"> The  smartest app to organize your fight </p>
        </div>
        <div>
          <f7-button
                  class="header__button_add"
                  @click="isCreateCharacter = true"> ☺
          </f7-button>
          <f7-button
                  class="header__button_add"
                  @click="isCreateCharacterList = true"> +
          </f7-button>
        </div>
      </header>

      <f7-block>
        <div v-if="listOfFightCharactersData.length">
          <div v-for="character in listOfFightCharactersData"
               :key="character.ID">
            <f7-button v-on:click="toFightList(character)"  style="background-color: #f0e6d5; margin: 5px; min-height: 140px; color: saddlebrown">
              <div>
                <h3> {{ character.Name }} id: {{ character.ID }} <br>
                  ArmorType: {{ character.ArmorType }} <br>
                  Initiative: {{ character.Initiative }} <br>
                  Hits: {{ character.Hits }}</h3>
              </div>
            </f7-button>
          </div>
        </div>
      </f7-block>

      <create-character-popup
        :opened="isCreateCharacter"
        @popup:complete="saveCharacter"
        @popup:closed="isCreateCharacter = false">
      </create-character-popup>

      <create-battle-list-popup
        :opened="isCreateCharacterList"
        :listOfCharactersData="getListOfCharacters"
        :listOfFightCharactersData="getListOfFightCharacters"
        @popup:complete="saveCharacter"
        @popup:closed="isCreateCharacterList = false">
      </create-battle-list-popup>

    </div>
  </f7-app>
</template>

<script>
  import { mapGetters } from "vuex";
  import CreateCharacterPopup from "./CreateCharacterPopup"
  import CreateBattleListPopup from "./CreateBattleListPopup"

  export default {
    name: 'app',
    components: {
      CreateCharacterPopup,
      CreateBattleListPopup
    },
    data() {
      return {
        isCreateCharacter: false,
        isCreateCharacterList: false,
        listOfCharactersData: [],
        listOfFightCharactersData: []
      }
    },
    computed: {
      ...mapGetters({
        getListOfCharacters: "characters/getListOfCharacters",
        getListOfFightCharacters: "characters/getListOfFightCharacters"
      }),
    },
    methods: {
      saveCharacter(data) {
        this.listOfFightCharactersData = data.slice();
      },
      toFightList(character) {
        let index = this.listOfFightCharactersData.findIndex((char) => char.ID === character.ID);
        this.listOfFightCharactersData = this.listOfFightCharactersData.filter((item,ind) => ind !== index);
        this.listOfCharacters.push(character);
      },
    },
    created() {
      this.$store.dispatch("characters/loadListOfCharacters");
    }
  }
</script>

<style>
  @import url('https://fonts.googleapis.com/css?family=Cinzel:700&display=swap');
  .room {
    background-image: url("../assets/background-dragon.jpg");
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    text-align: center;
    width: 100%;
    min-height: 1080px;
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
  }
  .header {
    display: flex;
    flex-direction: row;
    justify-content: center;
    color: white;
    text-align: left;
    margin: 1em;
    margin-top: 1.5em;
  }
  .header__text-title {
    font-family: 'Cinzel', serif;
    font-weight: lighter;
    letter-spacing: 2px;
    line-height: 110%;
    font-size: 35px;
    padding: 0;
    margin: 0 1em 0 0;
  }
  .header__image {
    color: white;
    width: 40px;
  }
  .header__button_add {
    font-size: 45px;
    padding: 0;
    margin: 0;
    height: 45px;
    width: 45px;
    color: white;
  }
  .card {
    background-color: #B9A487;
    margin: 5px;
  }
</style>