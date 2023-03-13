import { FaSolidSpinner } from 'solid-icons/fa';
import type { Component } from 'solid-js';

const Loading: Component = () => {
  return (
    <span>
      <span class="icon">
        <FaSolidSpinner />
      </span>
      <span>Loading...</span>
    </span>
  );
};

export default Loading;
