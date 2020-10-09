<template>
    <q-page padding>
        <div class="blog-edit-container">
            <message-banner :message="message" :color="messageBannerColor" v-if="showMessage">
                <template v-slot:action>
                    <q-btn flat color="white" label="GO BACK" @click="closeMessageBoxAndGoBack" />
                </template>
            </message-banner>
            <q-form @submit="submitBlog" class="q-gutter-md">
                <q-input v-model="blogView.blog.title" label="Title" />
                <q-select
                        v-model="blogView.type"
                        :options="typeList"
                        label="Type"
                        option-label="name">
                </q-select>
                <q-select
                        v-model="blogView.tags"
                        label="Tags"
                        use-input
                        use-chips
                        multiple
                        input-debounce="0"
                        :options="tagList"
                        option-label="name">
                </q-select>
                <q-input v-model="blogView.blog.imageUrl" label="URL of the front image"></q-input>
                <div class="q-pa-md">
                    <mavon-editor
                            v-model="blogView.blog.content"
                            :ishljs="true"
                            :toolbars="markdownOption"/>
                </div>
                <div>
                    <q-toggle label="Publish" v-model="blogView.blog.published" />
                </div>
                <q-btn label="Submit" type="submit" color="primary"/>
            </q-form>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    import MessageBanner from 'components/MessageBanner.vue'
    export default {
        name: 'BlogEdit',
        components: {
            MessageBanner
        },
        data () {
            return {
                blogId: '',
                blogView: {
                    blog: {
                        title: '',
                        type_id: '',
                        content: '',
                        imageUrl: '',
                        published: false
                    },
                    type: {
                        type_id: '',
                        name: ''
                    },
                    tags: []
                },
                typeList: [],
                tagList: [],
                markdownOption: {
                    bold: true, // 粗体
                    italic: true, // 斜体
                    header: true, // 标题
                    underline: true, // 下划线
                    strikethrough: true, // 中划线
                    mark: true, // 标记
                    superscript: true, // 上角标
                    subscript: true, // 下角标
                    quote: true, // 引用
                    ol: true, // 有序列表
                    ul: true, // 无序列表
                    code: true, // code
                    table: true, // 表格
                    subfield: true, // 单双栏模式
                    preview: true // 预览
                },
                // message banner related data
                message: '',
                messageBannerColor: '',
                showMessage: false,
                timer: null
            }
        },
        methods: {
            getBlogById () {
                axios.get('/api/blogs/' + this.blogId)
                    .then(this.getBlogByIdSuccess)
            },
            getBlogByIdSuccess (res) {
                if (res.data) {
                    const blogView = res.data
                    this.blog = blogView.blog
                    this.type = blogView.type
                    this.selectedTags = blogView.tags
                }
            },
            listTypesSuccess (res) {
                if (res.data) {
                    this.typeList = res.data
                }
            },
            listTagsSuccess (res) {
                if (res.data) {
                    this.tagList = res.data
                }
            },
            submitBlog () {
                this.blogView.blog.update_at = new Date().toISOString()
                this.blogView.blog.type_id = this.blogView.type.type_id
                if (this.blogId === null || this.blogId === undefined || this.blogId === '') {
                    // create new blog
                    this.blogView.blog.create_at = new Date().toISOString()
                    axios.post('/api/blogs', this.blogView)
                        .then(this.createBlogSuccess)
                } else {
                    // edit existing blog
                    axios.put('/api/blogs/', this.blogView)
                        .then(this.editBlogSuccess)
                }
            },
            createBlogSuccess () {
                this.setMessageAndColor('Successfully created the blog, going back in 3 seconds...', this.consts.MESSAGE_BOX_COLOR_POSITIVE)
                this.setMessageBoxTimer(3000, this.closeMessageBoxAndGoBack)
            },
            editBlogSuccess () {
                this.setMessageAndColor('Successfully edited the blog, going back in 3 seconds...', this.consts.MESSAGE_BOX_COLOR_POSITIVE)
                this.setMessageBoxTimer(3000, this.closeMessageBoxAndGoBack)
            },
            closeMessageBoxAndGoBack () {
                this.showMessage = false
                this.$router.push('/admin/blogs/')
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
        },
        mounted () {
            axios.get('/api/types')
                .then(this.listTypesSuccess)
            axios.get('/api/tags')
                .then(this.listTagsSuccess)
            this.blogId = this.$route.params.blog_id
            if (this.blogId !== null && this.blogId !== undefined && this.blogId !== '') {
                // edit existing blog
                this.getBlogById()
            }
        }
    }
</script>

<style scoped>
    .blog-edit-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
