import ShadowDOM from 'react-shadow';
import parse, { domToReact } from 'html-react-parser';
import type { DOMNode } from 'html-dom-parser';
import type { PhishingSimulationRecognitionFeatureValue } from '@api/index';
import { Avatar, Card, Popover, PopoverSurface, PopoverTrigger } from '@fluentui/react-components';
import { DismissRegular, SquareMultipleRegular, SubtractRegular } from '@fluentui/react-icons';

import MailDetailMockStyles from './MailDetailMock.module.scss';
import MailDetailMockShadowStyles from './MailDetailMock.shadow.scss?inline';

type MailDetailMockProps = {
  sentAt: string;
  sender: string;
  subject: string;
  content: string;
  recognitionFeatures: Array<PhishingSimulationRecognitionFeatureValue>;
}

export default function MailDetailMock({ sentAt, sender, subject, content, recognitionFeatures }: MailDetailMockProps) {
  const userLocale = navigator.language || 'de-DE';
  const sentAtFormatter = new Intl.DateTimeFormat(userLocale, {
    weekday: 'short',
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false
  });
  const sentDateTime = new Date(sentAt);

  const popoverValueDict: PopoverValueDict = {};
  recognitionFeatures.forEach(el => popoverValueDict[el.id] = (el.educationalInstruction ?? el.generalEducationalInstruction));
  
  
  return (
    <div className={MailDetailMockStyles.MailDetailMock}>
      <div className={MailDetailMockStyles.MailDetailMock__WindowHead}>
      <p className={MailDetailMockStyles.MailDetailMock__WindowHeadTitle}>
        { subject }
      </p>
      <div className={MailDetailMockStyles.MailDetailMock__WindowActions}>
        <SubtractRegular fontSize="16px" />
        <SquareMultipleRegular fontSize="16px" />
        <DismissRegular fontSize="16px" />
      </div>
    </div>
    <div className={MailDetailMockStyles.MailDetailMock__WindowBody}>
      <Card>
        { subject }
      </Card>
      <Card>
        <div className={MailDetailMockStyles.MailDetailMock__MailHeader}>
          <Avatar color="colorful" name={sender} size={36} />
          <div className={MailDetailMockStyles.MailDetailMock__MailHeaderDetails}>
            <div className={MailDetailMockStyles.MailDetailMock__MailHeaderRow}>
              <p className={MailDetailMockStyles.MailDetailMock__MailSenderName}>
                { sender }
              </p>
              <div className={MailDetailMockStyles.MailDetailMock__MailHeaderActions}>

              </div>
            </div>
            <div className={MailDetailMockStyles.MailDetailMock__MailHeaderRow}>
              <p className={MailDetailMockStyles.MailDetailMock__MailRecipientText}>An: Sie</p>
              <p>
                { sentAtFormatter.format(sentDateTime) }
              </p>
            </div>
          </div>
        </div>
        <div
          onClickCapture={(event) => {
            event.preventDefault();
            event.stopPropagation();
          }}
        >
          <HTMLWithPopovers htmlString={content} popoverValueDict={popoverValueDict} />
        </div>
      </Card>
    </div>
    </div>
  )
}

type PopoverValueDict = {
  [key: string]: string;
}

type HTMLWithPopoversProps = {
  htmlString: string;
  popoverValueDict: PopoverValueDict
};

function HTMLWithPopovers({ htmlString, popoverValueDict }: HTMLWithPopoversProps) {
  const extractedParts = extractPartsFromHTML(htmlString);

  const options = {
    replace: (domNode: DOMNode) => {
      if (
        domNode.type === 'tag' &&
        domNode.name === 'span' &&
        domNode.attribs &&
        domNode.attribs['data-feature-value-id']
      ) {
        const featureValueId = domNode.attribs['data-feature-value-id'];
        const educationalInstruction = popoverValueDict[featureValueId];
        if (educationalInstruction) {
          const spanProps = domNode.attribs;
          const children = domToReact((domNode.children as DOMNode[]), options);

          return (
            <Popover openOnHover withArrow>
              <PopoverTrigger>
                <span {...spanProps}>{children}</span>
              </PopoverTrigger>
              <PopoverSurface tabIndex={-1} className={MailDetailMockStyles.MailDetailMock__Popover}>
                { educationalInstruction }
              </PopoverSurface>
            </Popover>
          );
        }
      }
    },
  };

  return (
    <ShadowDOM.div>
      <style>
        {MailDetailMockShadowStyles}
        {extractedParts.styles}
      </style>
      <div className="MailDetailMock__Shadow">
        { parse(extractedParts.body, options) }
      </div>
    </ShadowDOM.div>
  );
};

function extractPartsFromHTML(html: string): { body: string, styles: string } {
  const parser = new DOMParser();
  const doc = parser.parseFromString(html, 'text/html');
  const styleTags = doc.querySelectorAll('style');
  let styles = '';
  styleTags.forEach(styleTag => {
    styles += styleTag.textContent;
  });
  return { body: doc.body.innerHTML, styles };
}