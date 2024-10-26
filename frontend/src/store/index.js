import { setBlockTracking } from 'vue';
import { createStore } from 'vuex';

var stores =  createStore({
  state: {
    email: '',
    password: '',
    username: '',
    bio: '',
    terms: '2024-09-03'
  },
  getters: {
    getEmail: state => state.email,
    getPassword: state => state.password,
    getUsername: state => state.username,
    getBio: state => state.bio,
    getTerms: state => state.terms,
  },
  mutations: {
    setEmail (state, email) {
        state.email = email
      },
      setPassword (state, password) {
        state.password = password
      },
      setUsername (state, username) {
        state.username = username
      },
      setBio (state, bio) {
        state.bio = bio
      }
  }
});

export default stores;