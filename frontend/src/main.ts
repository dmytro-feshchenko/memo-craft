import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import './configs/themes/axolotl-light.css';
// import useTheme from './modules/themes/utils/useTheme';

// useTheme();

import { bind } from './events'
bind();

createApp(App).mount('#app')
