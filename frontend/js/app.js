import Vue from 'vue';

import ShortcutList from './components/ShortcutList.vue';

Vue.component('shortcut-list', ShortcutList)

new Vue({
    el: '#app',
    components: {
        ShortcutList
    }
})