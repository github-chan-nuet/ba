import { DismissRegular, SquareMultipleRegular, SubtractRegular } from '@fluentui/react-icons';
import MailDetailMockStyles from './MailDetailMock.module.scss';
import { Avatar, Card, Popover, PopoverSurface, PopoverTrigger, tokens } from '@fluentui/react-components';
import parse, { domToReact } from 'html-react-parser';
import type { DOMNode } from 'html-dom-parser';

type MailDetailMockProps = {
  sentAt: string;
  sender: string;
  subject: string;
  content: string;
}

export default function MailDetailMock({ sentAt, sender, subject, content }: MailDetailMockProps) {
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
              <p
                className={MailDetailMockStyles.MailDetailMock__MailSenderName}
                style={{
                  color: tokens.colorBrandForeground1
                }}
              >
                { sender }
              </p>
              <div className={MailDetailMockStyles.MailDetailMock__MailHeaderActions}>

              </div>
            </div>
            <div
              className={MailDetailMockStyles.MailDetailMock__MailHeaderRow}
              style={{
                color: tokens.colorNeutralForeground3
              }}
            >
              <p className={MailDetailMockStyles.MailDetailMock__MailRecipientText}>An: Sie</p>
              <p>
                { sentAtFormatter.format(sentDateTime) }
              </p>
            </div>
          </div>
        </div>
        <div
          className={MailDetailMockStyles.MailDetailMock__MailBody}
          onClickCapture={(event) => {
            event.preventDefault();
            event.stopPropagation();
          }}
        >
          <HTMLWithPopovers htmlString={content} />
        </div>
      </Card>
    </div>
    </div>
  )
}

type HTMLWithPopoversProps = {
  htmlString: string;
};

function HTMLWithPopovers({ htmlString }: HTMLWithPopoversProps) {
  const bodyOnlyHTML = extractBodyHTML(htmlString);

  const options = {
    replace: (domNode: DOMNode) => {
      if (
        domNode.type === 'tag' &&
        domNode.name === 'span' &&
        domNode.attribs &&
        domNode.attribs['data-recogfeatid']
      ) {
        const spanProps = domNode.attribs;
        const children = domToReact((domNode.children as DOMNode[]), options);

        return (
          <Popover>
            <PopoverTrigger>
              <span {...spanProps}>{children}</span>
            </PopoverTrigger>
            <PopoverSurface tabIndex={-1}>
              Dies ist ein Beispiel-Inhalt
            </PopoverSurface>
          </Popover>
        );
      }
    },
  };

  return <div>{ parse(bodyOnlyHTML, options) }</div>
};

function extractBodyHTML(html: string): string {
  const parser = new DOMParser();
  const doc = parser.parseFromString(html, 'text/html');
  return doc.body.innerHTML;
}