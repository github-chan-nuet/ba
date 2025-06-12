import CommonGroundLesson1 from "@data/components/CommonGround/Lesson1";
import CommonGroundLesson2 from "@data/components/CommonGround/Lesson2";
import CommonGroundLesson3 from "@data/components/CommonGround/Lesson3";
import CommonGroundLesson4 from "@data/components/CommonGround/Lesson4";
import CommonGroundLesson5 from "@data/components/CommonGround/Lesson5";

import AngriffsvektorenLesson1 from "@data/components/Angriffsvektoren/Lesson1";
import AngriffsvektorenLesson2 from "@data/components/Angriffsvektoren/Lesson2";
import AngriffsvektorenLesson3 from "@data/components/Angriffsvektoren/Lesson3";

export type CourseRecord = {
  id: string;
  handle: string;
  label: string;
  description: string;
  lessons: LessonRecord[];
};

export type LessonRecord = {
  id: string;
  handle: string;
  label: string;
  description: string;
  contentElement: JSX.Element;
};

export async function getCourses() {
  return staticCourses.getAll();
}

export async function getCourse(handle: string) {
  return staticCourses.get(handle);
}

const staticCourses = {
  async getAll(): Promise<CourseRecord[]> {
    return staticCourses.records;
  },

  async get(handle: string): Promise<CourseRecord | null> {
    return staticCourses.records.find(course => course.handle === handle) || null;
  },

  records: [
    {
      id: '95850db0-7b1a-4ba1-863e-46b0ad066241',
      handle: 'common-ground',
      label: 'Common Ground',
      description: 'Dieser Kurs bietet einen kompakten Einstieg in das Thema Phishing. Du erfährst, was Phishing ist, wie verbreitet es ist, welche Schäden entstehen können und warum es für Securaware relevant ist.',
      lessons: [
        {
          id: '647ec003-43f2-46c3-ac31-891df5cb6d51',
          handle: 'extent-of-damage',
          label: 'Schadenausmass von Phishing',
          description: 'Lerne, wie gross der Schaden durch Phishing sein kann - sowohl finanziell als auch auf persönlicher und unternehmerischer Ebene - und warum Prävention so wichtig ist.',
          contentElement: <CommonGroundLesson1 />
        },
        {
          id: '57a1c0f1-49fc-43af-9100-a3a8b133d1db',
          handle: 'definition',
          label: 'Definition von Phishing',
          description: 'Erfahre, was Phishing ist, wie Angreifer vorgehen und woran du typische Merkmale wie Täuschung, Dringlichkeit und gefälschte Links erkennst.',
          contentElement: <CommonGroundLesson2 />
        },
        {
          id: 'dbce4c4e-2d81-4ccc-ac35-5fbd01f56164',
          handle: 'commonness',
          label: 'Verbreitung',
          description: 'Verstehe, warum Phishing so weit verbreitet ist und welche Faktoren wie globale Reichweite, menschliches Verhalten oder geringe Einstiegshürden es begünstigen.',
          contentElement: <CommonGroundLesson3 />
        },
        {
          id: 'f346691a-59c9-4b98-89dd-b27d52a872aa',
          handle: 'consequences-and-examples',
          label: 'Konsequenzen und Beispiele',
          description: 'Lerne, welche Auswirkungen ein Phishing-Angriff haben kann - von Geldverlust bis Identitätsdiebstahl - und entdecke reale Beispiele, die grosse Folgen hatten.',
          contentElement: <CommonGroundLesson4 />
        },
        {
          id: '40737135-e1e2-475a-bd02-890f34dd49b4',
          handle: 'motivation-for-securaware',
          label: 'Motivation für Securaware',
          description: 'Finde heraus, warum es sich lohnt, bei Securaware mitzumachen und aktiv zu beleiben - ob zum Schutz deiner Daten, aus beruflicher Verantwortung oder um deine Familie und Gesellschaft zu schützen.',
          contentElement: <CommonGroundLesson5 />
        }
      ]
    },
    {
      id: 'db0c1481-238f-4402-be2c-6d8efaabc0fc',
      handle: 'angriffsvektoren',
      label: 'Angriffsvektoren',
      description: 'Lerne, über welche Wege Phishing-Angriffe verbreitet werden - von E-Mail über SMS bis hin zu Telefonanrufen. Dieser Kurs zeigt dir die typischen Einfallstore für Angreifer.',
      lessons: [
        {
          id: '598b08fe-b5d9-47e7-b709-99b6047ca523',
          handle: 'email',
          label: 'E-Mail',
          description: 'Lerne, wie Cyberkriminelle E-Mails nutzen, um mit täuschend echten Nachrichten persönliche Daten zu stehlen. Erfahre, welche typischen Merkmale solche Angriffe haben und wie du verdächtige Inhalte zuverlässig entlarven kannst.',
          contentElement: <AngriffsvektorenLesson1 />
        },
        {
          id: '4f619101-c678-4807-85a4-2dbc9c5a7114',
          handle: 'sms',
          label: 'SMS',
          description: 'Entdecke die typischen Merkmale von Smishing-Angriffen, warum diese besonders tückisch sind und wie du gefährliche Nachrichten erkennst, bevor Schaden entsteht. Ideal für den sicheren Umgang mit mobilen Geräten.',
          contentElement: <AngriffsvektorenLesson2 />
        },
        {
          id: '31728ecf-2577-4492-bef6-6c7404444403',
          handle: 'telephone-call',
          label: 'Telefonanruf',
          description: 'Lerne typische Maschen und psychologische Tricks von Betrügern kennen, die dich am Telefon zur Preisgabe sensibler Informationen bringen wollen - und erfahre, wie du sicher reagierst.',
          contentElement: <AngriffsvektorenLesson3 />
        }
      ]
    },
  ] as Array<CourseRecord>,
}