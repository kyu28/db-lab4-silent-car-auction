import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue';
import ElementPlus from 'element-plus';
import './style.css';
import 'element-plus/dist/index.css';
import 'element-plus/theme-chalk/dark/css-vars.css';

const baseURL = "http://localhost:8000";
const tables = {
    "auto": {
        table: "auto",
        immutableCols: [],
        mutableCols: ['location', 'year', 'type_of_auto', 'mileage', 'VIN'],
        baseURL: baseURL
    },
    "auto-index": {
        table: "auto-index",
        immutableCols: [ "auto_id" ],
        mutableCols: [ 'VIN' ],
        baseURL: baseURL
    },
    "customer": {
        table: "customer",
        immutableCols: [],
        mutableCols: [ "first_name", "last_name", "location", "email" ],
        baseURL: baseURL
    },
    "customer-index": {
        table: "customer-index",
        immutableCols: [ "customer_id" ],
        mutableCols: [ "email" ],
        baseURL: baseURL
    },
    "transaction": {
        table: "transaction",
        immutableCols: [ "transaction_id" ],
        mutableCols: [ "auto_id", "customer_id", "bid", "bid_date" ],
        baseURL: baseURL
    },
    "winner": {
        table: "winner",
        immutableCols: [],
        mutableCols: [ "auto_id", "customer_id" ],
        baseURL: baseURL
    }
}

const lists = {
    "ford": {
        list: "Available Fords",
        cols: [ "auto_id", "location", "year", "type_of_auto", "mileage" ],
        url: baseURL + "/auto/FORD",
        hasArgs: false
    },
    "cheverolet": {
        list: "Available Cheverolets",
        cols: [ "auto_id", "location", "year", "type_of_auto", "mileage" ],
        url: baseURL + "/auto/CHEVEROLET",
        hasArgs: false
    },
    "search-autos-by-brand": {
        list: "Search autos by brand",
        cols: [ "auto_id", "location", "year", "type_of_auto", "mileage" ],
        url: baseURL + "/auto",
        hasArgs: true
    },
    "max-bid": {
        list: "Max Bid",
        cols: [ "auto_id", "location", "year", "maximum_bid" ],
        url: baseURL + "/transaction/max-bid",
        hasArgs: false
    },
    "winners-and-losers": {
        list: "Winners and Losers",
        cols: [ "auto_id", "last_name", "bid", "maximum_bid", "is_winner" ],
        url: baseURL + "/customer/winners-and-losers",
        hasArgs: true
    }
}

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/table/:table',
            component: () => import('./components/Table.vue'),
            props: route => tables[route.params.table]
        },
        {
            path: '/list/:list',
            component: () => import('./components/List.vue'),
            props: route => lists[route.params.list]
        },
        {
            path: '/',
            component: () => import('./components/Index.vue')
        },
        {
            path: '/login',
            component: () => import('./components/Login.vue')
        },
        {
            path: '/signup',
            component: () => import('./components/Signup.vue'),
            props: { baseURL: baseURL }
        },
        {
            path: '/transaction/new',
            component: () => import('./components/NewBid.vue'),
            props: { url: baseURL + "/transaction" }
        }
    ]
});
    
createApp(App)
    .use(ElementPlus)
    .use(router)
    .mount('#app')
