<template>
    <!-- Popup -->
    <f7-popup
            id="create-character-popup"
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
                <f7-block-title>Личные даннные</f7-block-title>
                <f7-list>
                    <f7-list-input
                            label="Имя"
                            floating-label
                            :value="name"
                            @blur="name = $event.target.value"
                            type="text"
                            placeholder="Введите имя">
                    </f7-list-input>
                    <f7-list-input
                            label="Тип доспеха"
                            floating-label
                            :value="armorType"
                            @blur="armorType = $event.target.value"
                            type="text"
                            placeholder="Введите тип доспеха">
                    </f7-list-input>
                    <f7-list-input
                            label="Инициатива"
                            floating-label
                            :value="initiative"
                            @blur="initiative = $event.target.value"
                            type="text"
                            placeholder="Введите инициативу">
                    </f7-list-input>
                    <f7-list-input
                            label="Количество хитов"
                            floating-label
                            :value="hits"
                            @blur="hits = $event.target.value"
                            type="text"
                            placeholder="Введите количество хитов">
                    </f7-list-input>
                </f7-list>
            </f7-block>
            </f7-page>
    </f7-popup>
</template>

<script>
    export default {
        name: "CreateCharacterPopup",
        props: {
            opened: {
                type: Boolean,
                default: false
            },
        },
        data() {
            return {
                name: "",
                armorType: "",
                initiative: null,
                hits: null,
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
        }
    }
</script>

<style src="../css/popup.css"></style>