import { useParams } from '@solidjs/router';
import { Component } from 'solid-js';
import Hero from '../layouts/Hero';

const Room: Component = () => {
  const { id } = useParams<{id: string}>();

  return (
    <Hero>
      <p>{id}</p>
    </Hero>
  );
};

export default Room;
