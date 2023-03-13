import { type Component, type JSX, Match, Switch, Resource, children } from 'solid-js';
import ApiRes from '../models/ApiRes';

const Res: Component<{children: JSX.Element, data: Resource<ApiRes<any>>}> = (props) => {
  const c = children(() => props.children);

  const { data } = props;

  return (
    <Switch fallback={<div>Not Found</div>}>
    <Match when={data.state === 'pending' || data.state === 'unresolved'}>
      Loading...
    </Match>
    <Match when={data.state === 'ready'}>
      {c()}
    </Match>
    <Match when={data.state === 'errored'}>
      {JSON.stringify(data.error)}
    </Match>
  </Switch>
  );
};

export default Res;
