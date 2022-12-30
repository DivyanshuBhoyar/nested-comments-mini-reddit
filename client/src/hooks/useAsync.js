import { useCallback, useEffect, useState } from "react";

export function useAsync(func, dependencies = []) {
  const { execute, ...state } = useAsyncInternal(func, dependencies, true); // initial loading is set true

  useEffect(() => {
    execute();
  }, [execute]);
  return state;
}

// function version -- return execute function to the caller to manually call
export function useAsyncFnc(func, dependencies = []) {
  return useAsyncInternal(func, dependencies, false);
}

const useAsyncInternal = (func, dependencies, initialLoading = false) => {
  const [loading, setLoading] = useState(initialLoading);
  const [error, setError] = useState(false);
  const [value, setValue] = useState();

  const exec = useCallback(({ ...params }) => {
    setLoading(true);
    return func(...params)
      .then((data) => {
        setValue(data);
        setError(undefined);
        return data;
      })
      .catch((e) => {
        setValue(undefined);
        setError(e);
        return Promise.reject(error);
      })
      .finally(() => setLoading(false));
  }, dependencies);

  return { loading, error, value, execute };
};
