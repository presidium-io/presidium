/*
 * Presidio
 *
 * Context aware, pluggable and customizable PII anonymization service for text and images.
 *
 * The version of the OpenAPI document: 2.0
 * Contact: presidio@microsoft.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using Microsoft.Presidio.Api;
using Microsoft.Presidio.Model;
using Microsoft.Presidio.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace Microsoft.Presidio.Test.Model
{
    /// <summary>
    ///  Class for testing AnalyzeRequest
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class AnalyzeRequestTests : IDisposable
    {

        private AnalyzeRequest instance;

        public AnalyzeRequestTests()
        {
            // TODO uncomment below to create an instance of AnalyzeRequest
            //instance = new AnalyzeRequest();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of AnalyzeRequest
        /// </summary>
        [Test]
        public void AnalyzeRequestInstanceTest()
        {
            Assert.IsInstanceOfType(typeof(AnalyzeRequest), instance, "variable 'instance' is a AnalyzeRequest");
        }


        /// <summary>
        /// Test the property 'Text'
        /// </summary>
        [Test]
        public void TextTest()
        {
            Assert.IsTrue(instance.Text == "hello world");
        }
        /// <summary>
        /// Test the property 'Language'
        /// </summary>
        [Test]
        public void LanguageTest()
        {
            Assert.IsTrue(instance.Language == "en");
        }
        /// <summary>
        /// Test the property 'CorrelationId'
        /// </summary>
        [Test]
        public void CorrelationIdTest()
        {
            Assert.IsTrue(instance.CorrelationId == "1234");
        }
        /// <summary>
        /// Test the property 'ScoreThreshold'
        /// </summary>
        [Test]
        public void ScoreThresholdTest()
        {
            Assert.IsTrue(instance.ScoreThreshold == 0.4);
        }
        /// <summary>
        /// Test the property 'Entities'
        /// </summary>
        [Test]
        public void EntitiesTest()
        {
            Assert.IsTrue(instance.Entities.Count.Equals(1));
        }
        /// <summary>
        /// Test the property 'ReturnDecisionProcess'
        /// </summary>
        [Test]
        public void ReturnDecisionProcessTest()
        {
            Assert.IsFalse(instance.ReturnDecisionProcess.Value);
        }


    }

}