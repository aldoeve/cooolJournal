import './assets/main.css';

import router from './routes/index';
import stores from './store';
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

import InputText from 'primevue/inputtext';
import Checkbox from 'primevue/checkbox';
import Button from 'primevue/button';
import Password from 'primevue/password';
import Textarea from 'primevue/textarea';

import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);


app.component('InputText', InputText);
app.component('Checkbox', Checkbox);
app.component('Button', Button);
app.component('Password', Password);
app.component('Textarea', Textarea);

app.use(router).use(PrimeVue, { theme: { preset: Aura } }).use(stores).mount('#app');
