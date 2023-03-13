import { children } from 'solid-js';
import type { Component, JSX } from 'solid-js';

const Hero: Component<{children: JSX.Element}> = (props) => {
  const c = children(() => props.children);

  return (
    <section class="hero is-fullheight-with-navbar">
      <div class="hero-body">
        {c()}
      </div>
    </section>
  );
};

export default Hero;
