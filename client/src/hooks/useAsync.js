import { useCallback, useEffect, useState } from "react";

// execute function will auto called initially
export function useAsync(func, dependencies = []) {
  const { exec, ...state } = useAsyncInternal(func, dependencies, true); // initial loading is set true

  // 1. Change in dependencies updates the execute() callback function
  // 2. This updated function as a dependency to to useEffect will auto call it when updated. , eg change in ID
  // 3. implies it also auto run initially.
  useEffect(() => {
    exec();
  }, [exec]);

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

  const exec = useCallback((...params) => {
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

  // execute function will get updated on change on dependencies and get returned
  return { loading, error, value, exec };
};
