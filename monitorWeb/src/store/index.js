import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const state = {
    apis:[],
    androidID:''
}

const getters = {
    get_apis(state){
        return state.apis
    },
    getAndroidID(state){
        return state.androidID
    }
}

const mutations ={
    set_apis(state,value){
        state.apis = value

    },
    setAndroidID(state,value){
        state.androidID = value
    }
}
export default new Vuex.Store(
    {
        getters,
        state,
        mutations
    }
)