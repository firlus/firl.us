<template>
    <div>
        <li v-for="shortcut in list" :key="shortcut.Path">
            {{ shortcut.Path }}: {{ shortcut.URL}}
        </li>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                list: null,
            }    
        },
        mounted() {
            axios
                .get('https://firl.us/api/shortcut')
                .then(response => (this.list = response.data))
        }
    }
</script>
