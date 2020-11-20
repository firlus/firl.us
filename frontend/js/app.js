import Vue from 'vue';
import 'bootstrap';

import ShortcutList from './components/ShortcutList.vue';
import ShortcutCreator from './components/ShortcutCreator.vue';
import Shortcut from './components/Shortcut.vue';

Vue.component('shortcut-list', ShortcutList);
Vue.component('shortcut-creator', ShortcutCreator);
Vue.component('shortcut', Shortcut);

new Vue({
    el: '#app',
    components: {
        ShortcutList,
        Shortcut
    }
})