import { observable, decorate } from "mobx";
import { createContext } from "react";

class AppStore {
  devices = {};
  isInitialLoadComplete = false;
  networkMetrics = {};
  appVersion = {};
}

decorate(AppStore, {
  devices: observable,
  isInitialLoadComplete: observable,
  networkMetrics: observable,
  appVersion: observable,
});

export const AppStoreContext = createContext(new AppStore());
