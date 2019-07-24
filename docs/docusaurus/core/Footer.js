/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

class Footer extends React.Component {
  docUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    const docsUrl = this.props.config.docsUrl;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    return `${baseUrl}${docsPart}${langPart}${doc}`;
  }

  pageUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    return baseUrl + (language ? `${language}/` : '') + doc;
  }

  render() {
    return (
      <footer className="nav-footer" id="footer">
        <section className="sitemap">
          <a href={this.props.config.baseUrl} className="nav-home">
            {this.props.config.footerIcon && (
              <img
                src={this.props.config.baseUrl + this.props.config.footerIcon}
                alt={this.props.config.title}
                width="66"
                height="58"
              />
            )}
          </a>
          <div>
            <h5>Docs</h5>
            <a href="https://github.com/facebookincubator/magma/blob/master/docs/Magma_Product_Overview.pdf">
              Magma Product Overview
            </a>
            <a href="https://github.com/facebookincubator/magma/blob/master/docs/Magma_Specs_V1.1.pdf">
              Magma Spec
            </a>
            <a href="https://github.com/facebookincubator/magma/blob/master/docs/Magma_Deployment_Guide_Installing_Magma_V2.pdf">
              Magma Install Guide
            </a>
            <a href="https://github.com/facebookincubator/magma/blob/master/docs/Magma_Network_Management_System.pdf">
              Magma Network Management System
            </a>
          </div>
          <div>
            <h5>Community</h5>
            <a href="https://discord.gg/4YxZbft">
              Discord
            </a>
            <a href="https://groups.google.com/forum/#!forum/magma-dev">
              Forum
            </a>
            <a
              href="http://stackoverflow.com/questions/tagged/"
              target="_blank"
              rel="noreferrer noopener">
              Stack Overflow
            </a>
            <a
              href="https://facebookincubator.github.io/magma/docs/devsummit"
              target="_blank"
              rel="noreferrer noopener">
              Dev Summit
            </a>
          </div>
          <div>
            <h5>More</h5>
            <a href="https://code.fb.com/open-source/magma/">Blog</a>
            <a href="https://github.com/facebookincubator/magma">GitHub</a>
          </div>
        </section>

        <a
          href="https://code.facebook.com/projects/"
          target="_blank"
          rel="noreferrer noopener"
          className="fbOpenSource">
          <img
            src={`${this.props.config.baseUrl}img/oss_logo.png`}
            alt="Facebook Open Source"
            width="170"
            height="45"
          />
        </a>
        <section className="copyright">{this.props.config.copyright}</section>
      </footer>
    );
  }
}

module.exports = Footer;
