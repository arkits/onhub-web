import { observable, decorate } from "mobx";
import { createContext } from "react";

class AppStore {
  devices = {};
  isInitialLoadComplete = false;
}

decorate(AppStore, {
  devices: observable,
  isInitialLoadComplete: observable,
});

export const AppStoreContext = createContext(new AppStore());
