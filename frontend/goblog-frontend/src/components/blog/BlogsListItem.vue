<template>
    <q-card class="blog-list-item-card q-mx-auto shadow-5 cursor-pointer" @click="handleCardClick">
        <q-card-section horizontal>
            <q-img ratio="4/3" v-if="blog.image_url !== ''" class="col-3 blog-img" :src="blog.image_url" alt="" />
            <q-card-section class="col">
                <!-- blog title and type area -->
                <q-card-section class="text-h5" horizontal>
                    <div class="full-width row">
                        <q-card-section class="col-9">
                            {{blog.title}}
                        </q-card-section>
                        <q-card-section class="col-3">
                            <q-chip outline color="secondary" text-color="white" icon="bookmark">
                                {{type.name}}
                            </q-chip>
                        </q-card-section>
                    </div>
                </q-card-section>
                <!-- blog content area -->
                <q-card-section class="text-body1">
                    <div class=blog-content>
                        {{blog.content}}
                    </div>
                </q-card-section>
                <!-- tags of this blog -->
                <q-card-section horizontal>
                    <q-card-section class="col-8">
                        <q-chip outline color="deep-purple-4" text-color="white" icon="label"
                                v-for="tag in tags" :key="tag.tag_id">
                            {{tag.name}}
                        </q-chip>
                    </q-card-section>
                    <q-card-section class="col-4 text-body2 text-right blog-date">
                        {{date}}
                    </q-card-section>
                </q-card-section>
            </q-card-section>
        </q-card-section>
    </q-card>
</template>

<script>
    export default {
        name: 'BlogsListItem',
        props: {
            blog: Object,
            type: Object,
            tags: Array
        },
        data () {
            return {
                date: String
            }
        },
        methods: {
            handleCardClick () {
                this.$router.push('/blogs/' + this.blog.blog_id)
            }
        },
        mounted () {
            // convert date format
            this.date = new Date(this.blog.create_at).toLocaleDateString()
        }
    }
</script>

<style scoped>
    .blog-list-item-card {
        width: 100%;
        max-width: 1000px;
        max-height: 500px;
        overflow: hidden;
        text-decoration: none;
    }
    .blog-content {
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 6;
        -webkit-box-orient: vertical;
    }
    .blog-img {
        max-height: 100%;
    }
    .blog-date {
        margin: auto;
    }
</style>
