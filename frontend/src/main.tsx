import ReactDOM from 'react-dom/client';
import { HashRouter, Route, Routes } from "react-router-dom";
import Helper from './about/Helper';
import App from './App';
import './main.css';
import Option from './options/Option';
ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <HashRouter basename={"/"}>
    {/* The rest of your app goes here */}
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/option" element={<Option />} />
      <Route path="/helper" element={<Helper />} />
      {/* <Route path="/page1" element={<Page1 />} />
      <Route path="/page2" element={<Page2 />} /> */}
      {/* more... */}
    </Routes>
  </HashRouter>,
)