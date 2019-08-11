<template>
    <div>
        <h3> {{ listOfCharacters }} </h3>
        <div>
            <p> {{ listOfCharacters.ArmorType }}</p>
            <p> {{ listOfCharacters.Initiative }}</p>
            <p> {{ listOfCharacters.Hits }}</p>
        </div>
    </div>
</template>

<script>
    export default {
        name: "CharacterCard",
        props: {
            characterData: {
                type: Object,
                default: {
                    ArmorType: "",
                    BattleID: null,
                    Hits: null,
                    ID: null,
                    Initiative: null,
                    Name: "",
                    UserID: null,
                }
            },
        },
        data() {
            return {
                // name: this.listOfCharacters.Name || "",
                // armorType: this.listOfCharacters.ArmorType || "",
                // initiative: this.listOfCharacters.Initiative || null,
                // hits: this.listOfCharacters.Hits || null,
                listOfCharacters: {
                    ArmorType: "",
                    BattleID: null,
                    Hits: null,
                    ID: null,
                    Initiative: null,
                    Name: "",
                    UserID: null,
                },
                listOfFightCharacters: this.listOfFightCharactersData || []
            }
        },
        watch: {
            characterData(v) {
                this.listOfCharacters = {
                    ArmorType: v.ArmorType,
                    BattleID: v.BattleID,
                    Hits: v.Hits,
                    ID: v.ID,
                    Initiative: v.Initiative,
                    Name: v.Name,
                    UserID: v.UserID,
                };
                console.log(v);
                v.forEach((item) => console.log(item));
            }
        }
    }
</script>

<style scoped>

</style>