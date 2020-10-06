<template>
    <q-page padding>
        <div class="types-admin-container">
            <message-banner :message="message" :showMessage="showMessage" />
            <q-list bordered class="rounded-borders shadow-3">
                <q-item-label header>Create New Type</q-item-label>
                <q-item>
                    <q-input
                            style="width: 100%"
                            v-model="newTypeName"
                            label="Type Name"
                            hint="1~20 characters"
                            :rules="[ typeName => typeName.length > 0 && typeName.length <= 20 || '1~20 characters']">
                        <template v-slot:append>
                            <q-btn flat dense round size="12px" icon="add" @click="createNewType"></q-btn>
                        </template>
                    </q-input>

                </q-item>
                <q-separator spaced/>
                <q-item-label header>Types</q-item-label>
                <q-item v-for="type in typeList" :key="type.type_id">
                    <q-item-section>
                        {{type.name}}
                    </q-item-section>
                    <q-item-section side>
                        <div>
                            <q-btn flat dense round size="12px" icon="edit">
                                <q-popup-edit
                                        v-model="type.name"
                                        :validate="typeName => {return typeName.length > 0 && typeName.length <= 20}"
                                        @save="editTypeName(type)">
                                    <q-input
                                            v-model="type.name"
                                            dense
                                            autofocus
                                            counter
                                            hint="1~20 chars"></q-input>
                                </q-popup-edit>
                            </q-btn>
                            <q-btn flat dense round size="12px" icon="delete" @click="deleteType(type)"></q-btn>
                        </div>
                    </q-item-section>
                </q-item>
            </q-list>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    import MessageBanner from 'components/MessageBanner.vue'
    export default {
        name: 'TypesAdmin',
        components: {
            MessageBanner
        },
        data () {
            return {
                newTypeName: '',
                typeList: [],
                message: '',
                showMessage: false,
                timer: null
            }
        },
        methods: {
            refreshTypeList () {
                axios.get('/api/types')
                    .then(this.listTypesSuccess)
            },
            listTypesSuccess (res) {
                if (res.data) {
                    this.typeList = res.data
                }
            },
            createNewType () {
                // check that newTypeName is valid

                axios.post('/api/types', {
                    name: this.newTypeName
                }).then(() => {
                    this.newTypeName = ''
                    this.setMessageAndRefresh('Successfully created the type')
                })
            },
            editTypeName (type) {
                axios.put('/api/types', type)
                    .then(this.refreshTypeList)
            },
            deleteType (type) {
                axios.delete('/api/types/' + type.type_id)
                    .then(() => {
                        this.setMessageAndRefresh('Successfully deleted the type')
                    })
            },
            setMessageAndRefresh (msg) {
                this.message = msg
                this.showMessage = true
                this.refreshTypeList()
                window.clearTimeout(this.timer)
                this.timer = window.setTimeout(() => {this.showMessage = false}, 3000)
            }
        },
        mounted () {
            this.refreshTypeList()
        }
    }
</script>

<style scoped>
    .types-admin-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
