<template>
    <transition
        appear
        enter-active-class="animated fadeIn"
        leave-active-class="animated fadeOut"
        >
        <q-banner
            class="text-white shadow-4 absolute-top message-banner"
            :class="color"
            rounded
            dense
            inline-actions
            style="z-index: 99">
            {{message}}
            <template v-slot:action>
                <q-btn flat color="white" label="GOT IT" @click="closeMessageBox" />
            </template>
        </q-banner>
    </transition>
</template>

<script>
export default {
    name: 'MessageBanner',
    props: {
        message: {
            type: String,
            required: true
        },
        color: {
            type: String
        }
    },
    methods: {
        closeMessageBox () {
            this.$emit('close')
        }
    }
}
</script>

<style scoped>
    .message-banner {
        max-width: 500px;
        margin: 5px auto;
    }
</style>