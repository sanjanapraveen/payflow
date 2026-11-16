import { App } from './components';
import { initializeLib } from './lib';

const startApp = () => {
    initializeLib();
    const app = new App();
    app.render();
};

startApp();