import MainLayout from 'layouts/MainLayout'
import Blogs from 'pages/blog/Blogs'
import BlogDetail from 'pages/blog/BlogDetail'
import Types from 'pages/Types'
import Tags from 'pages/Tags'
import About from 'pages/About'
import Archives from 'pages/Archives'
import TypesAdmin from 'pages/admin/TypesAdmin'
import TagsAdmin from 'pages/admin/TagsAdmin'

const routes = [
    {
        path: '/',
        component: MainLayout,
        children: [{
            path: '/blogs',
            component: Blogs
        }, {
            path: '/blogs/:id',
            component: BlogDetail
        }, {
            path: '/types',
            component: Types
        }, {
            path: '/tags',
            component: Tags
        }, {
            path: '/about',
            component: About
        }, {
            path: '/archives',
            component: Archives
        }, {
            path: '/admin/types',
            component: TypesAdmin
        }, {
            path: '/admin/tags',
            component: TagsAdmin
        }
        ]
    },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
