import { Link, useNavigate } from '@solidjs/router';
import { type Component, createResource, createSignal, For, Show } from 'solid-js';
import { api } from '../common';
import Loading from '../components/Loading';
import Hero from '../layouts/Hero';

const Rooms: Component = () => {
  const navigate = useNavigate();
  const [rooms] = createResource(() => api.getRooms());
  const [createRoom, setCreateRoom] = createSignal('');

  const startCreateRoom = async () => {
    const res = await api.createRoom(createRoom())

    if (!res.success) {
      alert("Error")
      return
    }

    navigate(`/rooms/${res.data.slug}`)
  }

  return (
    <Hero>
      <div class="container has-text-centered">
        <p class="title">Rooms</p>
        <Show when={!rooms.loading} fallback={<Loading />}>
          <>
            {rooms()!.data.length === 0 && <p>No rooms available!</p>}
            <div class="columns is-centered is-vcentered is-mobile">
              <For each={rooms()!.data}>{(room, i) =>
                <div class="column is-narrow">
                  <div class="card">
                    <div class="card-content">
                      <p>{room.name}</p>
                      <p>{room.description}</p>
                    </div>
                    <footer class="card-footer">
                      <Link class="card-footer-item" href={`/rooms/${room.slug}`}>Go</Link>
                    </footer>
                  </div>
                </div>
              }</For>
            </div>
          </>
        </Show>
        <div class="columns is-centered is-vcentered">
          <div class="column is-one-quarter">
            <div class="field has-addons has-addons-centered">
              <div class="control">
                <input class="input" type="text" onChange={(e) => setCreateRoom(e.currentTarget.value)} />
              </div>
              <div class="control">
                <button class="button is-success" onClick={startCreateRoom}>New room</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Hero>
  );
};

export default Rooms;
