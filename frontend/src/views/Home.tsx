import { type Component, createSignal } from 'solid-js';
import Hero from '../layouts/Hero';
import { FaSolidUser } from 'solid-icons/fa';
import { api } from '../common';
import Storage from '../helpers/Storage';
import { A, useNavigate } from '@solidjs/router';

const Home: Component = () => {
  const navigate = useNavigate();
  const [username, setUsername] = createSignal('');

  const startLogin = async () => {
    const token = await api.login(username());
    if (token === "") {
      alert("Error")
      return
    }

    api.setAuth(token, username());
    Storage.set("token", token);
    Storage.set("username", username());
    navigate("/rooms");
  }

  return (
    <Hero>
      <div class="container has-text-centered">
        <div class="columns is-centered is-vcentered">
          <div class="column is-narrow">
            <figure class="image is-128x128 is-inline-block">
              <img class="is-rounded" src="/assets/coding.jpg" />
            </figure>
          </div>
          <div class="column is-narrow">
            <p class="title">Are ya' coding, Son?</p>
            <p class="subtitle">Code editing in real time</p>
          </div>
        </div>
        <div class="columns is-centered is-vcentered">
          <div class="column is-one-quarter">
            {api.isLoggedIn() ? (
              <>
                <p>{api.getUsername()}</p>
                <A href="/rooms" class="button is-success">Go</A>
              </>
            ) : (
              <div class="field has-addons has-addons-centered">
                <div class="control">
                  <input class="input" type="text" placeholder="Your usuername" onChange={(e) => setUsername(e.currentTarget.value)} />
                </div>
                <div class="control">
                  <button class="button is-success" onClick={startLogin}>Go</button>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </Hero>
  );
};

export default Home;
