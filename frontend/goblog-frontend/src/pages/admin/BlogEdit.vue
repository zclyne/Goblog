<template>
    <q-page padding>
        <div class="blog-edit-container">
            <q-banner class="bg-primary text-white" inline-actions v-if="msg">
                {{msg}}
                <template v-slot:action>
                    <q-btn flat color="white" label="Go Back" @click="backToBlogsAdmin" />
                </template>
            </q-banner>
            <q-form @submit="submitBlog" class="q-gutter-md">
                <q-input v-model="blog.title" label="Title" />
                <q-select
                        v-model="type"
                        :options="typeList"
                        label="Type"
                        :option-label="(type) => type.name !== undefined ? type.name : ''">
                </q-select>
                <q-select
                        v-model="selectedTags"
                        label="Tags"
                        use-input
                        use-chips
                        multiple
                        input-debounce="0"
                        :options="tagList"
                        option-label="name">
                </q-select>
                <q-input v-model="blog.imageUrl" label="URL of the front image"></q-input>
                <div class="q-pa-md">
                    <mavon-editor
                            v-model="blog.content"
                            :ishljs="true"
                            :toolbars="markdownOption"/>
                </div>
                <div>
                    <q-toggle label="Publish" v-model="blog.published" />
                </div>
                <q-btn label="Submit" type="submit" color="primary"/>
            </q-form>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'BlogEdit',
        data () {
            return {
                blogId: '',
                blog: {},
                title: '',
                type: {},
                selectedTags: [],
                content: '',
                imageUrl: '',
                published: false,
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
                    /* 2.2.1 */
                    subfield: true, // 单双栏模式
                    preview: true // 预览
                },
                msg: ''
            }
        },
        methods: {
            getBlogById () {
                axios.get('/api/blogs/' + this.blogId)
                    .then(this.getBlogByIdSuccess)
            },
            getBlogByIdSuccess (res) {
                console.log(res)
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
                this.blog.update_at = new Date().toISOString()
                const blogView = {
                    blog: this.blog,
                    type: this.type,
                    tags: this.selectedTags
                }
                if (this.blogId === null || this.blogId === undefined || this.blogId === '') {
                    // create new blog
                    blogView.blog.create_at = new Date().toISOString()
                    axios.post('/api/blogs', blogView)
                        .then(this.createBlogSuccess)
                } else {
                    // edit existing blog
                    axios.put('/api/blogs/', blogView)
                        .then(this.editBlogSuccess)
                }
            },
            createBlogSuccess () {
                this.msg = 'Successfully created a new blog'
            },
            editBlogSuccess () {
                this.msg = 'Successfully edited the blog'
            },
            backToBlogsAdmin () {
                this.$router.push('/admin/blogs/')
            }
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
