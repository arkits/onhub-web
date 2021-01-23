import { observable, decorate } from "mobx";
import { createContext } from "react";

class AppStore {
  devices = {};
  isInitialLoadComplete = false;
  networkMetrics = {};
}

decorate(AppStore, {
  devices: observable,
  isInitialLoadComplete: observable,
  networkMetrics: observable,
});

export const AppStoreContext = createContext(new AppStore());
