<template>
    <q-page padding>
        <div class="column blog-detail-container q-mt-md q-mx-auto">
            <message-banner :message="message" :color="messageBannerColor" v-if="showMessage">
                <template v-slot:action>
                    <q-btn flat color="white" label="OK" @click="closeMessageBox" />
                </template>
            </message-banner>
            <div class="row">
                <div class="col-9">
                    <div class="text-h4">
                        {{blog.title}}
                    </div>
                    <div class="text-grey-9 q-my-md col-auto">
                        {{date}}
                    </div>
                </div>
                <div class="col-3">
                    <q-chip outline color="secondary float-right" text-color="white" icon="bookmark">
                        {{type.name}}
                    </q-chip>
                </div>
            </div>
            <div v-if="blog.image_url !== ''">
                <q-img :src="blog.image_url"></q-img>
            </div>
            <div>
                <mavon-editor
                        style="width: 100%"
                        :editable="false"
                        :subfield="false"
                        defaultOpen="preview"
                        :toolbarsFlag="false"
                        :scrollStyle="true"
                        :ishljs="true"
                        :boxShadow="false"
                        :value="blog.content" />
            </div>
            <div class="row q-pt-md">
                <q-chip outline color="deep-purple-4"
                        text-color="white"
                        icon="label"
                        v-for="tag in tags" :key="tag.tag_id">
                    {{tag.name}}
                </q-chip>
            </div>

            <!-- comment area -->
            <q-list bordered class="q-mt-md rounded-borders">
                <div v-for="(comment, index) in comments" :key="comment.comment_id">
                    <q-item class="q-pa-md">
                        <q-item-section>
                            <q-item-label>{{comment.nickname}}</q-item-label>
                            <q-item-label caption>{{comment.content}}</q-item-label>
                        </q-item-section>
                        <q-item-section side top>
                            <q-item-label caption>{{new Date(comment.create_at).toLocaleString()}}</q-item-label>
                        </q-item-section>
                    </q-item>
                    <q-separator v-if="index < comments.length - 1" />
                </div>
                <q-separator spaced v-if="comments.length > 0"/>
                <!-- add comment area -->
                <div class="q-mt-md q-mx-md">
                    <div class="text-h6">Leave Your Comment Here</div>
                    <q-input label="Nickname" v-model="commenterNickname"></q-input>
                    <q-input label="Email" v-model="commenterEmail"></q-input>
                    <q-input label="Your Thoughts" type="textarea" v-model="commentContent"></q-input>
                    <q-btn color="primary" label="Submit" class="float-right q-my-md" @click="submitCommentBtnClick"/>
                </div>
            </q-list>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    import MessageBanner from 'components/MessageBanner.vue'
    export default {
        name: 'BlogDetail',
        components: {
            MessageBanner
        },
        data () {
            return {
                blog: '',
                type: '',
                tags: [],
                date: '',
                comments: [],
                commenterNickname: '',
                commenterEmail: '',
                commentContent: '',
                // message banner related data
                message: '',
                messageBannerColor: '',
                showMessage: false,
                timer: null
            }
        },
        methods: {
            getBlogByIdSuccess (res) {
                if (res.data) {
                    this.blog = res.data.blog
                    this.type = res.data.type
                    this.tags = res.data.tags
                    // convert date format
                    this.date = new Date(this.blog.create_at).toLocaleDateString()
                }
            },
            listCommentsByBlogIdSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.comments = res.data
                }
            },
            submitCommentBtnClick () {
                if (this.commenterNickname === '') {
                    this.setMessageAndColor('Nickname should not be empty!', this.consts.MESSAGE_BOX_COLOR_WARNING)
                } else if (this.commenterEmail === '') {
                    this.setMessageAndColor('Email should not be empty!', this.consts.MESSAGE_BOX_COLOR_WARNING)
                } else if (this.commentContent === '') {
                    this.setMessageAndColor('Comment content should not be empty!', this.consts.MESSAGE_BOX_COLOR_WARNING)
                } else {
                    const submitCommentRequest = {
                        blog_id: this.blog.blog_id,
                        nickname: this.commenterNickname,
                        email: this.commenterEmail,
                        content: this.commentContent,
                        create_at: new Date()
                    }
                    axios.post('/api/comments', submitCommentRequest)
                    .then(this.refreshComments)
                    this.setMessageAndColor('Successfully added the comment', this.consts.MESSAGE_BOX_COLOR_POSITIVE)
                    this.setMessageBoxTimer(3000, this.closeMessageBox)
                }
            },
            refreshComments () {
                axios.get('/api/comments?blog_id=' + this.$route.params.blog_id)
                    .then(this.listCommentsByBlogIdSuccess)
                // clear input area
                this.commenterNickname = ''
                this.commenterEmail = ''
                this.commentContent = ''
            },
            setMessageAndColor (msg, color) {
                this.message = msg
                this.messageBannerColor = color
                this.showMessage = true
            },
            setMessageBoxTimer (timeInMilliSecond, func) {
                window.clearTimeout(this.timer)
                this.timer = window.setTimeout(func, timeInMilliSecond)
            },
            closeMessageBox () {
                this.showMessage = false
            }
        },
        mounted () {
            axios.get('/api/blogs/' + this.$route.params.blog_id)
                .then(this.getBlogByIdSuccess)
            this.refreshComments()
        }
    }
</script>

<style scoped>
    .blog-detail-container {
        max-width: 1000px;
    }
</style>
