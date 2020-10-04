<template>
    <q-layout view="lHh Lpr lFf">
        <q-header elevated>
            <q-toolbar>
                <q-btn
                        flat
                        dense
                        round
                        icon="menu"
                        aria-label="Menu"
                        @click="leftDrawerOpen = !leftDrawerOpen"
                />

                <q-toolbar-title>
                    Goblog Admin
                </q-toolbar-title>

                <q-btn flat label="LOGOUT" @click="logout" />
            </q-toolbar>
        </q-header>

        <q-drawer
                v-model="leftDrawerOpen"
                show-if-above
                bordered
                content-class="bg-grey-1"
        >
            <q-list style="height: calc(100% - 150px); margin-top: 150px; border-right: 1px solid #ddd">
                <EssentialLink
                        v-for="item in essentialLinks"
                        :key="item.title"
                        v-bind="item"
                />
            </q-list>
            <q-img class="absolute-top" src="https://cdn.quasar.dev/img/material.png" style="height: 150px">
                <div class="absolute-bottom bg-transparent">
                    <q-avatar size="56px" class="q-mb-sm">
                        <img src="../static/avatar.jpg">
                    </q-avatar>
                    <div class="text-weight-bold">Yifan Zhang</div>
                    <div>MSCS@Columbia University</div>
                </div>
            </q-img>
        </q-drawer>

        <q-page-container>
            <router-view />
        </q-page-container>
    </q-layout>
</template>

<script>
    import EssentialLink from 'components/EssentialLink.vue'

    const linksData = [
        {
            title: 'Blogs Admin',
            icon: 'article',
            path: '/admin/blogs'
        },
        {
            title: 'Types Admin',
            icon: 'bookmark',
            path: '/admin/types'
        },
        {
            title: 'Tags Admin',
            icon: 'label',
            path: '/admin/tags'
        }
    ]

    export default {
        name: 'AdminLayout',
        components: { EssentialLink },
        data () {
            return {
                leftDrawerOpen: true,
                essentialLinks: linksData
            }
        },
        methods: {
            logout () {
                this.$store.commit('user/deleteUserToken')
                this.$router.push('/')
            }
        }
    }
</script>
