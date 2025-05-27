import { Body2, Title2 } from "@fluentui/react-components";
import useAuth from "../utils/auth/useAuth";

import WelcomeBannerStyles from "../styles/WelcomeBanner.module.scss";

export default function WelcomeBanner() {
  const { user } = useAuth();

  return (
    <div className={WelcomeBannerStyles.WelcomeBanner}>
      <Title2>Sicherheit beginnt mit dir - bleib wachsam, {user?.firstname}!</Title2>
      <Body2 className={WelcomeBannerStyles.WelcomeBanner__text}>
        Wir helfen dir, dich gezielt vor Online-Bedrohungen zu schützen, indem wir dich unterstützen, Phishing und andere Cyberangriffe frühzeitig zu erkennen. Denn nur gemeinsam können wir für eine sichere Umgebung sorgen!
      </Body2>
      <div className={WelcomeBannerStyles.WelcomeBanner__illustrationContainer}>
        <img
          src={"/blob.svg"}
          className={WelcomeBannerStyles.WelcomeBanner__blob}
        />
        <img
          src={"/illustration_1.svg"}
          className={WelcomeBannerStyles.WelcomeBanner__ilustration}
        />
      </div>
    </div>
  );
}