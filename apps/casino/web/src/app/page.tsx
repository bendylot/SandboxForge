import Header from "../components/Header";
import Hero from "../components/Hero";
import { RegCurrency, FridayActions } from "../components/Sections";
import Games from "../components/Games";
import Roadmap from "../components/Roadmap";
import Footer from "../components/Footer";
import ModalSignup from "./(modals)/signup/ModalSignup";
import { getFlags } from "../lib/featureFlags";

export const revalidate = 60; // ISR: пересборка раз в минуту

export default function Page() {
  const flags = getFlags();
  return (
    <>
      <ModalSignup />
      <Header />
      <Hero />
      <RegCurrency />
      <FridayActions />
      {flags.showGames && <Games />}
      {flags.showRoadmap && <Roadmap />}
      <Footer />
    </>
  );
}
