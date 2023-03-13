import { Route, Routes } from '@solidjs/router';
import type { Component } from 'solid-js';
import { api } from './common';
import Navbar from './components/Navbar';
import Storage from './helpers/Storage';
import Home from './views/Home';
import Room from './views/Room';
import Rooms from './views/Rooms';

const App: Component = () => {
  const token = Storage.getToken()
  const username = Storage.getUsername()

  if (token !== "") {
    api.setAuth(token, username);
  }

  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" component={Home} />
        <Route path="/rooms" component={Rooms} />
        <Route path="/rooms/:id" component={Room} />
      </Routes>
    </>
  );
};

export default App;
